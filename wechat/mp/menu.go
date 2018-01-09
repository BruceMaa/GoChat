package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
)

const (
	WechatMenuCreateApi              = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s`         // 自定义菜单创建接口
	WechatMenuGetApi                 = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s`            // 自定义菜单查询接口
	WechatMenuDeleteApi              = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s`         // 自定义菜单删除接口
	WechatMenuConditionalAddApi      = `https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=%s` // 创建个性化菜单
	WechatMenuConditionalDeleteApi   = `https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=%s` // 删除个性化菜单
	WechatMenuConditionalTrymatchApi = `https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=%s`       // 测试个性化菜单匹配结果
)

type (
	Menu struct {
		Menuid string  `json:"menuid,omitempty"`
		Menu   Buttons `json:"menu,omitempty"`
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

	Conditional struct {
		Buttons
		Matchrule ConditionalRule `json:"matchrule"`
	}

	ConditionalRule struct {
		TagId              string        `json:"tag_id,omitempty"`                      // 用户标签的id，可通过用户标签管理接口获取
		Sex                WechatUserSex `json:"sex,string,omitempty"`                  // 性别：男（1）女（2），不填则不做匹配
		Country            string        `json:"country,omitempty"`                     // 国家信息，是用户在微信中设置的地区，具体请参考地区信息表
		Province           string        `json:"province,omitempty"`                    // 省份信息，是用户在微信中设置的地区，具体请参考地区信息表
		City               string        `json:"city,omitempty"`                        // 城市信息，是用户在微信中设置的地区，具体请参考地区信息表
		Language           string        `json:"language,omitempty"`                    // 语言信息，是用户在微信中设置的语言，具体请参考语言表： 1、简体中文 "zh_CN" 2、繁体中文TW "zh_TW" 3、繁体中文HK "zh_HK" 4、英文 "en" 5、印尼 "id" 6、马来 "ms" 7、西班牙 "es" 8、韩国 "ko" 9、意大利 "it" 10、日本 "ja" 11、波兰 "pl" 12、葡萄牙 "pt" 13、俄国 "ru" 14、泰文 "th" 15、越南 "vi" 16、阿拉伯语 "ar" 17、北印度 "hi" 18、希伯来 "he" 19、土耳其 "tr" 20、德语 "de" 21、法语 "fr"
		ClientPlatformType string        `json:"client_platform_type,string,omitempty"` // 客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配
	}
)

type WechatClientPlatformType int

const (
	WechatClientPlatformTypeIOS     WechatClientPlatformType = iota + 1 // iOS 系统
	WechatClientPlatformTypeAndroid                                     // Android 系统
	WechatClientPlatformTypeOthers                                      // 其他手机系统
)

// 自定义菜单创建接口
func (wm *WechatMp) CreateMenu(accessToken string, buttons Buttons) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WechatMenuCreateApi, accessToken), buttons)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu create error: %+v\n", err)
		return nil, fmt.Errorf("menu create error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu create error: %+v\n", err)
		return nil, fmt.Errorf("menu create error: %+v\n", err)
	}
	return &result, nil
}

// 自定义菜单查询接口
func (wm *WechatMp) GetMenu(accessToken string) (*Menu, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatMenuGetApi, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu get error: %+v\n", err)
		return nil, fmt.Errorf("menu get error: %+v\n", err)
	}
	var menu Menu
	if err = json.Unmarshal([]byte(resp), &menu); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu get error: %+v\n", err)
		return nil, fmt.Errorf("menu get error: %+v\n", err)
	}
	return &menu, nil
}

// 自定义菜单删除接口
func (wm *WechatMp) DeleteMenu(accessToken string) (*common.PublicError, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatMenuDeleteApi, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu delete error: %+v\n", err)
		return nil, fmt.Errorf("menu delete error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu delete error: %+v\n", err)
		return nil, fmt.Errorf("menu delete error: %+v\n", err)
	}
	return &result, nil
}

/* 自定义菜单事件推送 */

// 创建个性化菜单
func (wm *WechatMp) AddConditional(accessToken string, conditional Conditional) (string, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WechatMenuConditionalAddApi, accessToken), conditional)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu conditional add error: %+v\n", err)
		return "", fmt.Errorf("menu conditional add error: %+v\n", err)
	}
	var menu Menu
	if err = json.Unmarshal([]byte(resp), &menu); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu conditional add error: %+v\n", err)
		return "", fmt.Errorf("menu conditional add error: %+v\n", err)
	}
	return menu.Menuid, nil
}

// 删除个性化菜单
func (wm *WechatMp) DeleteConditional(accessToken string, menuid string) (*common.PublicError, error) {
	menu := &Menu{
		Menuid: menuid,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WechatMenuConditionalDeleteApi, accessToken), menu)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu conditional delete error: %+v\n", err)
		return nil, fmt.Errorf("menu conditional delete error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu conditional delete error: %+v\n", err)
		return nil, fmt.Errorf("menu conditional delete error: %+v\n", err)
	}
	return &result, nil
}

// 测试个性化菜单匹配结果
func (wm *WechatMp) TrymatchConditional(accessToken, userId string) (*Buttons, error) {
	var args = make(map[string]string)
	args["user_id"] = userId
	resp, err := common.HTTPPostJson(fmt.Sprintf(WechatMenuConditionalTrymatchApi, accessToken), args)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu conditional trymatch error: %+v\n", err)
		return nil, err
	}
	var buttons Buttons
	if err = json.Unmarshal([]byte(resp), buttons); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "menu conditional trymatch error: %+v\n", err)
		return nil, err
	}
	return &buttons, nil
}
