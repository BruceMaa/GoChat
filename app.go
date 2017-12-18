package main

import (
	"fmt"
	"github.com/golang/sync/errgroup"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	group errgroup.Group
)

func init() {
	// 配置日志
	//logFile, _ := os.Create("gochat.log")
	//errorLogFile, _ := os.Create("gochat_error.log")
	//DebugHandle = io.MultiWriter(logFile)
	//TraceHandle = io.MultiWriter(logFile)
	//InfoHandle = io.MultiWriter(logFile)
	//WarnHandle = io.MultiWriter(logFile)
	//ErrorHandle = io.MultiWriter(logFile)
	//
	//gin.DefaultWriter = io.MultiWriter(logFile)
	//gin.DefaultErrorWriter = io.MultiWriter(errorLogFile)
}

func main() {

	serverInner := &http.Server{
		Addr:         fmt.Sprintf(":%d", ChatConfig.InnerPort),
		Handler:      initRouterInner(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverOutter := &http.Server{
		Addr:         fmt.Sprintf(":%d", ChatConfig.OutterPort),
		Handler:      initRouterOutter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	Logger.Debug("当前运行环境:", gin.Mode())

	group.Go(func() error {
		return serverInner.ListenAndServe()
	})

	group.Go(func() error {
		return serverOutter.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		Logger.Errorf("run error: %+v", err)
	}
}

