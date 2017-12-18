package mp

import (
	"fmt"
	"testing"
)

func TestWechatMp_AccessToken(t *testing.T) {
	config := &WechatMpConfig{
		AppId:     "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token:     "bingobox",
	}
	wechatMp := &WechatMp{
		Configure: *config,
	}
	accessToken, err := wechatMp.AccessTokenFromWechat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
}
