package mp

import (
	"fmt"
	"testing"
)

var wechatMp_ips WechatMp
var token_ips string

func init() {
	config := &WechatMpConfig{
		AppId:     "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token:     "bingobox",
	}
	wechatMp_ips.Configure = *config
	accessToken, err := wechatMp_ips.AccessTokenFromWechat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
	token_ips = accessToken.AccessToken
}

func TestGetCallBackIP(t *testing.T) {
	callBackIP, err := wechatMp_ips.GetCallBackIP(token_ips)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", callBackIP)
}
