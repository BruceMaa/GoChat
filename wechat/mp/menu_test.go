package mp

import (
	"testing"
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
)

func TestWechatMp_CreateMenu(t *testing.T) {
	config := &WechatMpConfig{
		AppId: "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token: "bingobox",
	}
	wechatMp := &WechatMp{
		Configure: *config,
	}
	accessToken, err := wechatMp.AccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)

	buttons := &Buttons{}

	button1 := &Button{}
	button1.Name = "扫码"
	subButton11 := &Button{}
	subButton11.Name = "扫码带提示"
	subButton11.Type = "scancode_waitmsg"
	subButton11.Key = "rselfmenu_0_0"
	button1.SubButton = append(button1.SubButton, *subButton11)
	subButton12 := &Button{}
	subButton12.Name = "扫码推事件"
	subButton12.Type = "scancode_push"
	subButton12.Key = "rselfmenu_0_1"
	button1.SubButton = append(button1.SubButton, *subButton12)
	buttons.Button = append(buttons.Button, *button1)

	button2 := &Button{}
	button2.Name = "发图"
	subButton21 := &Button{}
	subButton21.Name = "系统拍照发图"
	subButton21.Type = "pic_sysphoto"
	subButton21.Key = "rselfmenu_1_0"
	subButton22 := &Button{}
	subButton22.Name = "拍照或者相册发图"
	subButton22.Type = "pic_photo_or_album"
	subButton22.Key = "rselfmenu_1_1"
	subButton23 := &Button{}
	subButton23.Name = "微信相册发图"
	subButton23.Type = "pic_weixin"
	subButton23.Key = "rselfmenu_1_2"
	button2.SubButton = append(button2.SubButton, *subButton21)
	button2.SubButton = append(button2.SubButton, *subButton22)
	button2.SubButton = append(button2.SubButton, *subButton23)
	buttons.Button = append(buttons.Button, *button2)

	button3 := &Button{}
	button3.Name = "发送位置"
	button3.Type = "location_select"
	button3.Key = "rselfmenu_2_0"
	buttons.Button = append(buttons.Button, *button3)

	fmt.Printf("%+v\n", buttons)

	result, err := wechatMp.CreateMenu(accessToken.AccessToken, *buttons)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_GetMenu(t *testing.T) {
	config := &WechatMpConfig{
		AppId: "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token: "bingobox",
	}
	wechatMp := &WechatMp{
		Configure: *config,
	}
	accessToken, err := wechatMp.AccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)

	menu, err := wechatMp.GetMenu(accessToken.AccessToken)
	fmt.Printf("%+v\n", menu)
}

func TestWechatMp_DeleteMenu(t *testing.T) {
	config := &WechatMpConfig{
		AppId: "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token: "bingobox",
	}
	wechatMp := &WechatMp{
		Configure: *config,
	}
	accessToken, err := wechatMp.AccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)

	result, err := wechatMp.DeleteMenu(accessToken.AccessToken)
	fmt.Printf("%+v\n", result)
}

func TestSelfmenu(t *testing.T) {
	url := `https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s`
	config := &WechatMpConfig{
		AppId: "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token: "bingobox",
	}
	wechatMp := &WechatMp{
		Configure: *config,
	}
	accessToken, err := wechatMp.AccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)

	resp, err := common.HTTPGet(fmt.Sprintf(url, accessToken.AccessToken))
	fmt.Printf("%s\n", string(resp))
}
