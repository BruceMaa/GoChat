package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
)

const (
	WechatWebAuthCodeApi               = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s#wechat_redirect` // 用户同意授权，获取code
	WechatWebAuthAccessTokenApi        = `https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s`                                      // 通过code换取网页授权access_token
	WechatWebAuthRefreshAccessTokenApi = `https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=%s&refresh_token=%s`                                      // 刷新access_token（如果需要）
	WechatWebAuthUserinfoApi           = `https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s`                                                        // 拉取用户信息(需scope为 snsapi_userinfo)
	WechatWebAuthCheckApi              = `https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s`                                                                    // 检验授权凭证（access_token）是否有效
)

type WechatAuthScope string

const (
	WechatAuthScopeSnsapiBase     WechatAuthScope = "snsapi_base"     // 静默授权
	WechatAuthScopeSnsapiUserinfo WechatAuthScope = "snsapi_userinfo" // 弹出授权页面
)

type (
	// 微信网页授权accessToken
	WechatWebAuthAccessToken struct {
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		Openid       string `json:"openid"`        // 用户唯一标识
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
		common.PublicError
	}

	//微信网页授权获取用户信息
	WechatWebAuthUserInfo struct {
		Openid     string        `json:"openid"`     // 用户的唯一标识
		Nickname   string        `json:"nickname"`   // 用户昵称
		Sex        WechatUserSex `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
		Province   string        `json:"province"`   // 用户个人资料填写的省份
		City       string        `json:"city"`       // 普通用户个人资料填写的城市
		Country    string        `json:"country"`    // 国家，如中国为CN
		Headimgurl string        `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
		Privilege  []string      `json:"privilege"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
		Unionid    string        `json:"unionid"`    // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
		common.PublicError
	}
)

// 网页授权，拼接获取code的URL
func (wm *WechatMp) BuildWebAuthCodeURL(redirectUrl string, scope WechatAuthScope, state string) string {
	return fmt.Sprintf(WechatWebAuthCodeApi, wm.Configure.AppId, redirectUrl, "code", scope, state)
}

// 通过code换取网页授权access_token
func (wm *WechatMp) WebAuthAccessToken(code string) (*WechatWebAuthAccessToken, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatWebAuthAccessTokenApi, wm.Configure.AppId, wm.Configure.AppSecret, code, "authorization_code"))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthAccessToken http get error: %+v\n", err)
		return nil, err
	}
	var wechatWebAuthAccessToken WechatWebAuthAccessToken
	if err = json.Unmarshal(resp, &wechatWebAuthAccessToken); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthAccessToken json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &wechatWebAuthAccessToken, nil
}

// 刷新网页授权accessToken
func (wm *WechatMp) WebAuthRefreshAccessToken(refreshToken string) (*WechatWebAuthAccessToken, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatWebAuthRefreshAccessTokenApi, wm.Configure.AppId, "refresh_token", refreshToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthRefreshAccessToken http get error: %+v\n", err)
		return nil, err
	}
	var wechatWebAuthAccessToken WechatWebAuthAccessToken
	if err = json.Unmarshal(resp, &wechatWebAuthAccessToken); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthRefreshAccessToken json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &wechatWebAuthAccessToken, nil
}

// 网页授权，拉取用户信息(需scope为 snsapi_userinfo)
func (wm *WechatMp) WebAuthUserInfo(accessToken, openid string) (*WechatWebAuthUserInfo, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatWebAuthUserinfoApi, accessToken, openid, WechatLanguageZhCn))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthUserInfo http get error: %+v\n", err)
		return nil, err
	}
	var wechatWebAuthUserInfo WechatWebAuthUserInfo
	if err = json.Unmarshal(resp, &wechatWebAuthUserInfo); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthUserInfo json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &wechatWebAuthUserInfo, nil
}

// 检验授权凭证（access_token）是否有效
func (wm *WechatMp) WebAuthCheck(accessToken, openid string) (*common.PublicError, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatWebAuthCheckApi, accessToken, openid))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthCheck http get error: %+v\n", err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "WebAuthCheck json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &result, nil
}
