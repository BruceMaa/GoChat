package mp

import (
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
	"time"
)

type (
	AccessTokenHandlerFunc                func(wm *WechatMp) string
	SubscribeHandlerFunc                  func(request *MsgEventSubscribe) string
	UnSubscribeHandlerFunc                func(request *MsgEventSubscribe) string
	ScanHandlerFunc                       func(request *MsgEventScan) string
	LocationHandlerFunc                   func(request *MsgEventLocation) string
	MenuClickHandlerFunc                  func(request *MsgEventMenuClick) string
	MenuViewHandlerFunc                   func(request *MsgEventMenuView) string
	QualificationVerifySuccessHandlerFunc func(request *MsgEventQualificationVerifySuccess) string
	QualificationVerifyFailHandlerFunc    func(request *MsgEventQualificationVerifyFail) string
	NamingVerifySuccessHandlerFunc        func(request *MsgEventNamingVerifySuccess) string
	NamingVerifyFailHandlerFunc           func(request *MsgEventNamingVerifyFail) string
	AnnualRenewHandlerFunc                func(request *MsgEventAnnualRenew) string
	VerifyExpireHandlerFunc               func(request *MsgEventVerifyExpired) string
	SendTemplateFinishHandlerFunc         func(request *MsgEventTemplateSendFinish) string
	TextMessageHandlerFunc                func(request *TextMessage) string
	ImageMessageHandlerFunc               func(request *ImageMessage) string
	VoiceMessageHandlerFunc               func(request *VoiceMessage) string
	VideoMessageHandlerFunc               func(request *VideoMessage) string
	ShortVideoMessageHandlerFunc          func(request *ShortVideoMessage) string
	LocationMessageHandlerFunc            func(request *LocationMessage) string
	LinkMessageHandlerFunc                func(request *LinkMessage) string
)

// 默认获取微信accessToken的方法
func WechatMpDefaultAccessTokenHandlerFunc(wm *WechatMp) string {
	// 如果有配置缓存redis

	// 没有配置redis，则保存在内存中
	if wm.AccessToken == nil {
		updateWechatMpAccessToken(wm)
	} else {
		now := time.Now().Unix()
		accessTokenLastUpdateTime := wm.AccessToken.lastUpdateTime
		expiredSeconds := wm.AccessToken.ExpiresIn
		if now > accessTokenLastUpdateTime+expiredSeconds {
			updateWechatMpAccessToken(wm)
		}
	}
	return wm.AccessToken.AccessToken
}

// 更新微信accessToken
func updateWechatMpAccessToken(wm *WechatMp) {
	accessToken, err := wm.AccessTokenFromWechat()
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "updateWechatMpAccessToken error: %+v\n", err)
	}
	wm.AccessToken = accessToken
}
