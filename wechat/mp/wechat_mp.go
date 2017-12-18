package mp

import (
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"io/ioutil"
	"net/http"
)

const (
	RESPONSE_STRING_SUCCESS = "success"
	RESPONSE_STRING_FAIL    = "fail"
	RESPONSE_STRING_INVALID = "invalid wechat server"

	WECHAT_LANGUAGE_ZH_CN = "zh_CH" // 微信语言，简体中文
	WECHAT_LANGUAGE_ZH_TW = "zh_TW" // 微信语言，繁体中文
	WECHAT_LANGUAGE_EN    = "en"    // 微信语言，英文
)

type (
	WechatMpConfig struct {
		AppId          string `json:"app_id"`                     // 公众号appId
		AppSecret      string `json:"app_secret"`                 // 公众号appSecret
		Token          string `json:"token"`                      // 公众号Token
		EncodingAESKey string `json:"encoding_aes_key,omitempty"` // 公众号EncodingAESKey
	}

	WechatMp struct {
		Configure                         WechatMpConfig
		AccessToken                       *WechatAccessToken                    // 保存微信accessToken
		AccessTokenHandler                AccessTokenHandlerFunc                // 处理微信accessToken，如果有缓存，可以将accessToken存储到缓存中，默认存储到内存中
		SubscribeHandler                  SubscribeHandlerFunc                  // 关注微信公众号处理方法
		UnSubscribeHandler                UnSubscribeHandlerFunc                // 取消关注公众号处理方法
		ScanHandler                       ScanHandlerFunc                       // 扫描此微信公众号生成的二维码处理方法
		LocationHandler                   LocationHandlerFunc                   // 上报地理位置的处理方法
		MenuClickHandler                  MenuClickHandlerFunc                  // 自定义菜单点击的处理方法
		MenuViewHandler                   MenuViewHandlerFunc                   // 自定义菜单跳转外链的处理方法
		QualificationVerifySuccessHandler QualificationVerifySuccessHandlerFunc // 资质认证成功处理方法
		QualificationVerifyFailHandler    QualificationVerifyFailHandlerFunc    // 资质认证失败处理方法
		NamingVerifySuccessHandler        NamingVerifySuccessHandlerFunc        // 名称认证成功的处理方法
		NamingVerifyFailHandler           NamingVerifyFailHandlerFunc           // 名称认证失败的处理方法
		AnnualRenewHandler                AnnualRenewHandlerFunc                // 年审通知的处理方法
		VerifyExpiredHandler              VerifyExpireHandlerFunc               // 认证过期失效通知的处理方法
		SendTemplateFinishHandler         SendTemplateFinishHandlerFunc         // 发送模板消息结果通知
		//TextHandler        MsgTextHandlerFunc
		//ImageHandler       MsgImageHandlerFunc
		//VoiceHandler       MsgVoiceHandlerFunc
		//VideoHandler       MsgVideoHandlerFunc
	}
)

// 新建一个微信公众号
func New(wechatMpConfig *WechatMpConfig) *WechatMp {
	var wechatMp = &WechatMp{}
	wechatMp.Configure = *wechatMpConfig
	wechatMp.SetAccessTokenHandlerFunc(WechatMpDefaultAccessTokenHandlerFunc)
	return wechatMp
}

// 用户在设置微信公众号服务器配置，并开启后，微信会发送一次认证请求，此函数即做此验证用
func (wm *WechatMp) AuthWechatServer(r *http.Request) string {
	echostr := r.FormValue("echostr")
	if wm.checkWechatSource(r) {
		return echostr
	}
	return RESPONSE_STRING_INVALID
}

// 获取微信公众号的认证echo信息
func (wm *WechatMp) checkWechatSource(r *http.Request) bool {
	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")
	return CheckWechatAuthSign(signature, wm.Configure.Token, timestamp, nonce)
}

// 微信服务推送消息接收方法
func (wm *WechatMp) CallBackFunc(r *http.Request) string {
	// 首先，验证消息是否从微信服务发出
	valid := wm.checkWechatSource(r)
	if !valid {
		fmt.Fprintln(common.WechatErrorLoggerWriter, RESPONSE_STRING_INVALID)
		return RESPONSE_STRING_FAIL
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WechatRequest ioutil.ReadAll(body) error: %+v\n", err)
		return RESPONSE_STRING_FAIL
	}
	return wm.wechatMessageHandler(data)
}

// 设置全局获取微信accessToken的方法
func (wm *WechatMp) SetAccessTokenHandlerFunc(handlerFunc AccessTokenHandlerFunc) {
	wm.AccessTokenHandler = handlerFunc
}

// 设置处理关注事件的方法
func (wm *WechatMp) SetSubscribeHandlerFunc(handlerFunc SubscribeHandlerFunc) {
	wm.SubscribeHandler = handlerFunc
}

// 设置处理取消关注事件的方法
func (wm *WechatMp) SetUnSubscribeHandlerFunc(handlerFunc UnSubscribeHandlerFunc) {
	wm.UnSubscribeHandler = handlerFunc
}

// 设置处理扫描事件的方法
func (wm *WechatMp) SetScanHandlerFunc(handlerFunc ScanHandlerFunc) {
	wm.ScanHandler = handlerFunc
}

// 设置处理上报地理位置的方法
func (wm *WechatMp) SetLocationHandlerFunc(handlerFunc LocationHandlerFunc) {
	wm.LocationHandler = handlerFunc
}

// 设置处理自定义菜单点击事件的方法
func (wm *WechatMp) SetMenuClickHandlerFunc(handlerFunc MenuClickHandlerFunc) {
	wm.MenuClickHandler = handlerFunc
}

// 设置处理自定义菜单跳转外链事件的方法
func (wm *WechatMp) SetMenuViewHandlerFunc(handlerFunc MenuViewHandlerFunc) {
	wm.MenuViewHandler = handlerFunc
}

// 设置处理微信text消息事件方法
//func (wm *WechatMp) SetTextHandlerFunc(handlerFunc MsgTextHandlerFunc) {
//	wm.TextHandler = handlerFunc
//}

// 设置处理微信image消息事件方法
//func (wm *WechatMp) SetImageHandlerFunc(handlerFunc MsgImageHandlerFunc) {
//	wm.ImageHandler = handlerFunc
//}

// 设置处理微信voice消息事件方法
//func (wm *WechatMp) SetVoiceHandlerFunc(handlerFunc MsgVoiceHandlerFunc) {
//	wm.VoiceHandler = handlerFunc
//}

// 设置处理微信video消息事件方法
//func (wm *WechatMp) SetVideoHandlerFunc(handlerFunc MsgVideoHandlerFunc) {
//	wm.VideoHandler = handlerFunc
//}
