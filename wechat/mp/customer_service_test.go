package mp

import (
	"testing"
	"fmt"
	"time"
)

var token string
var wechatMp WechatMp
func init() {
	config := &WechatMpConfig{
		AppId: "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token: "bingobox",
	}
	wechatMp.Configure = *config
	accessToken, err := wechatMp.AccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
	token = accessToken.AccessToken
}

func TestWechatMp_GetKfList(t *testing.T) {
	customerServiceInfoList, _ := wechatMp.GetKfList(token)
	fmt.Printf("%+v\n", customerServiceInfoList)
}

func TestWechatMp_GetOnlineKfList(t *testing.T) {
	customerServiceInfoOnlineList, _ := wechatMp.GetOnlineKfList(token)
	fmt.Printf("%+v\n", customerServiceInfoOnlineList)
}

func TestWechatMp_AddKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
		Nickname: "test1",
	}
	result, _ := wechatMp.AddKfaccount(token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_InviteKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
		InviteWx: "test_kfwx",
	}
	result, _ := wechatMp.InviteKfaccount(token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_UpdateKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
		Nickname: "客服1",
	}
	result, _ := wechatMp.UpdateKfaccount(token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_UploadheadimgKfaccount(t *testing.T) {

}

func TestWechatMp_DeleteKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
	}
	result, _ := wechatMp.DeleteKfaccount(token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_CreateKfsession(t *testing.T) {
	var kfsession = &Kfsession{
		KfAccount: "test1@test",
		Openid: "openid",
	}
	result, _ := wechatMp.CreateKfsession(token, kfsession)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_CloseKfsession(t *testing.T) {
	var kfsession = &Kfsession{
		KfAccount: "test1@test",
		Openid: "openid",
	}
	result, _ := wechatMp.CloseKfsession(token, kfsession)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_GetKfsession(t *testing.T) {
	kfsession, _ := wechatMp.GetKfsession(token, "openid")
	fmt.Printf("%+v\n", kfsession)
}

func TestWechatMp_GetKfsessionList(t *testing.T) {
	kfsessionList, _ := wechatMp.GetKfsessionList(token, "openid")
	fmt.Printf("%+v\n", kfsessionList)
}

func TestWechatMp_GetWaitcaseList(t *testing.T) {
	waitcaseList, _ := wechatMp.GetWaitcaseList(token)
	fmt.Printf("%+v\n", waitcaseList)
}

func TestWechatMp_GetMsgrecordList(t *testing.T) {
	startTime := time.Now().Unix()
	after, _ := time.ParseDuration("1h")
	endTime := time.Now().Add(after).Unix()
	var param = &GetMsgListParam{
		Starttime: startTime,
		Endtime: endTime,
		Msgid: 1,
		Number: 10000,
	}
	getMsgListResp, _ := wechatMp.GetMsgrecordList(token, param)
	fmt.Printf("%+v\n", getMsgListResp)
}