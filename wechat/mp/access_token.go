package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
	"os"
	"time"
)

const (
	WECHAT_ACCESS_TOKEN_API = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s` // 获取微信accessToken地址
)

type (
	WechatAccessToken struct {
		AccessToken    string `json:"access_token"` // 公众号获取到的凭证
		ExpiresIn      int64  `json:"expires_in"`   // 公众号凭证有效时间，单位：秒
		lastUpdateTime int64
		*common.PublicError
	}
)

// 获取微信通行证
func (wm *WechatMp) AccessTokenFromWechat() (*WechatAccessToken, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_ACCESS_TOKEN_API, wm.Configure.AppId, wm.Configure.AppSecret))
	if err != nil {
		fmt.Fprintf(os.Stderr, "access_token get error: %+v\n", err)
		return nil, fmt.Errorf("access_token get error: %+v\n", err)
	}
	var wechatAccessToken WechatAccessToken
	if err = json.Unmarshal([]byte(resp), &wechatAccessToken); err != nil {
		fmt.Fprintf(os.Stderr, "access_token get error: %+v\n", err)
		return nil, fmt.Errorf("access_token get error: %+v\n", err)
	}
	wechatAccessToken.lastUpdateTime = time.Now().Unix()
	return &wechatAccessToken, nil
}
