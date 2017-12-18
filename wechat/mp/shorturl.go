package mp

import (
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"github.com/chanxuehong/wechat/json"
)

const WECHAT_SHORT_URL_API = `https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s`

const WECHAT_SHORT_URL_ACTION = "long2short"

type (
	ShortUrlRequest struct {
		Action  string `json:"action"`
		LongUrl string `json:"long_url"`
	}

	ShortUrlResponse struct {
		common.PublicError
		ShortUrl string `json:"short_url"`
	}
)

// 长链接转短链接接口
func (wm *WechatMp) Long2Short(accessToken, longUrl string) (*ShortUrlResponse, error) {
	shortUrlReq := &ShortUrlRequest{
		Action:  WECHAT_SHORT_URL_ACTION,
		LongUrl: longUrl,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_SHORT_URL_API, accessToken), shortUrlReq)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Long2Short common.HTTPPostJson error: %+v\n", err)
		return nil, err
	}

	var shortUrlResponse ShortUrlResponse
	if err = json.Unmarshal(resp, &shortUrlResponse); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Long2Short json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &shortUrlResponse, nil
}
