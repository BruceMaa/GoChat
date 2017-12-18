package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/mp"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type (
	ServerConf struct {
		Comment  string      `json:"comment"` // 配置文件说明
		ChatEnvs []AllConfig `json:"chat_envs"`
	}
	AllConfig struct {
		Mode       string     `json:"mode"`        // 当前运行环境
		ModeConfig ModeConfig `json:"mode_config"` // 当前运行环境的配置
	}
	ModeConfig struct {
		Name         string       `json:"name"`          // 运行环境名称
		OutterPort   int          `json:"outter_port"`   // 运行时对外网端口
		InnerPort    int          `json:"inner_port"`    // 运行时对内网端口
		WechatConfig WechatConfig `json:"wechat_config"` // 微信配置信息
	}
	WechatConfig struct {
		WechatMpConfig mp.WechatMpConfig `json:"wechat_mp_config"`
	}
)

// 配置文件地址
var configPath string

// 运行环境
var ginMode string

// 运行配置
var ChatConfig ModeConfig

// 初始化命令行参数
func init() {
	flag.StringVar(&configPath, "config", "/etc/bingobox/gochat.json", "config file of gochat")
	flag.StringVar(&ginMode, "mode", gin.DebugMode, "run mode of gochat")
	flag.Parse()

	extractConfig()
}

// 解析参数
func extractConfig() {
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open error: %+v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	serverConf := new(ServerConf)
	if err = decoder.Decode(&serverConf); err != nil {
		fmt.Println("error:", err)
		fmt.Fprintf(os.Stderr, "decoder.Decode() error: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("配置文件信息：%+v\n", serverConf)

	// 设置运行环境
	if strings.EqualFold(ginMode, gin.DebugMode) ||
		strings.EqualFold(ginMode, gin.TestMode) ||
		strings.EqualFold(ginMode, gin.ReleaseMode) {
		gin.SetMode(ginMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 根据运行环境，选择运行配置
	if len(serverConf.ChatEnvs) > 0 {
		for _, conf := range serverConf.ChatEnvs {
			if conf.Mode == gin.Mode() {
				ChatConfig = conf.ModeConfig
				break
			}
		}
	} else {
		// 配置文件有误
		fmt.Fprintln(os.Stderr, "配置文件有误")
		os.Exit(1)
	}

	fmt.Printf("当前配置信息%v\n", ChatConfig)
}
