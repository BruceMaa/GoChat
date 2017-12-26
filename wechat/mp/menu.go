package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
)

const (
	MENU_CREATE_API = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s` // 自定义菜单创建接口
	MENU_GET_API    = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s`    // 自定义菜单查询接口
	MENU_DELETE_API = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s` // 自定义菜单删除接口
)

type (
	Menu struct {
		Menu Buttons `json:"menu"`
	}

	Buttons struct {
		Button []Button `json:"button"` // 一级菜单数组，个数应为1~3个
	}

	Button struct {
		Name      string   `json:"name"`                 // 菜单标题，不超过16个字节，子菜单不超过60个字节
		Type      string   `json:"type,omitempty"`       // 菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
		Key       string   `json:"key,omitempty"`        // 菜单KEY值，用于消息接口推送，不超过128字节
		Url       string   `json:"url,omitempty"`        // 网页链接，用户点击菜单可打开链接，不超过1024字节。
		Appid     string   `json:"appid,omitempty"`      // 小程序的appid
		Pagepath  string   `json:"pagepath,omitempty"`   // 小程序的页面路径
		SubButton []Button `json:"sub_button,omitempty"` // 二级菜单数组，个数应为1~5个
	}
)

// 自定义菜单创建接口
func (wm *WechatMp) CreateMenu(accessToken string, buttons Buttons) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(MENU_CREATE_API, accessToken), buttons)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu create error: %+v\n", err)
		return nil, fmt.Errorf("menu create error: %+v\n", err)
	}
	var result common.PublicError
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu create error: %+v\n", err)
		return nil, fmt.Errorf("menu create error: %+v\n", err)
	}
	return &result, nil
}

// 自定义菜单查询接口
func (wm *WechatMp) GetMenu(accessToken string) (*Menu, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(MENU_GET_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu get error: %+v\n", err)
		return nil, fmt.Errorf("menu get error: %+v\n", err)
	}
	var menu Menu
	err = json.Unmarshal([]byte(resp), &menu)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu get error: %+v\n", err)
		return nil, fmt.Errorf("menu get error: %+v\n", err)
	}
	return &menu, nil
}

// 自定义菜单删除接口
func (wm *WechatMp) DeleteMenu(accessToken string) (*common.PublicError, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(MENU_DELETE_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu delete error: %+v\n", err)
		return nil, fmt.Errorf("menu delete error: %+v\n", err)
	}
	var result common.PublicError
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu delete error: %+v\n", err)
		return nil, fmt.Errorf("menu delete error: %+v\n", err)
	}
	return &result, nil
}

// TODO 自定义菜单事件推送
