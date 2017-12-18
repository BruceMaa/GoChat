package mp

import (
	"fmt"
	"testing"
)

var media_token string
var media_wechatMp WechatMp

func init() {
	config := &WechatMpConfig{
		AppId:     "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token:     "bingobox",
	}
	media_wechatMp.Configure = *config
	accessToken, err := media_wechatMp.AccessTokenFromWechat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
	media_token = accessToken.AccessToken
}
func TestWechatMp_GetMedia(t *testing.T) {
	var mediaId = "U_XJtS30wfLwcA9kohMwoNXX61kqIKnltVX7MgySs8AitRIF2kd_ryl5XBTILqyJ"
	media_wechatMp.GetMedia(media_token, mediaId)
}
