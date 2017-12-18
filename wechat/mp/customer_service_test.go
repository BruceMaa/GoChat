package mp

import (
	"fmt"
	"testing"
	"time"
)

var customer_service_token string
var customer_service_wechatMp WechatMp

func init() {
	config := &WechatMpConfig{
		AppId:     "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token:     "bingobox",
	}
	customer_service_wechatMp.Configure = *config
	accessToken, err := customer_service_wechatMp.AccessTokenFromWechat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
	customer_service_token = accessToken.AccessToken
}

func TestWechatMp_GetKfList(t *testing.T) {
	customerServiceInfoList, _ := customer_service_wechatMp.GetKfList(customer_service_token)
	fmt.Printf("%+v\n", customerServiceInfoList)
}

func TestWechatMp_GetOnlineKfList(t *testing.T) {
	customerServiceInfoOnlineList, _ := customer_service_wechatMp.GetOnlineKfList(customer_service_token)
	fmt.Printf("%+v\n", customerServiceInfoOnlineList)
}

func TestWechatMp_AddKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
		Nickname:  "test1",
	}
	result, _ := customer_service_wechatMp.AddKfaccount(customer_service_token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_InviteKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
		InviteWx:  "test_kfwx",
	}
	result, _ := customer_service_wechatMp.InviteKfaccount(customer_service_token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_UpdateKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
		Nickname:  "客服1",
	}
	result, _ := customer_service_wechatMp.UpdateKfaccount(customer_service_token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_UploadheadimgKfaccount(t *testing.T) {

}

func TestWechatMp_DeleteKfaccount(t *testing.T) {
	var kfInfo = &KfInfo{
		KfAccount: "test1@test",
	}
	result, _ := customer_service_wechatMp.DeleteKfaccount(customer_service_token, kfInfo)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_CreateKfsession(t *testing.T) {
	var kfsession = &Kfsession{
		KfAccount: "test1@test",
		Openid:    "openid",
	}
	result, _ := customer_service_wechatMp.CreateKfsession(customer_service_token, kfsession)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_CloseKfsession(t *testing.T) {
	var kfsession = &Kfsession{
		KfAccount: "test1@test",
		Openid:    "openid",
	}
	result, _ := customer_service_wechatMp.CloseKfsession(customer_service_token, kfsession)
	fmt.Printf("%+v\n", result)
}

func TestWechatMp_GetKfsession(t *testing.T) {
	kfsession, _ := customer_service_wechatMp.GetKfsession(customer_service_token, "openid")
	fmt.Printf("%+v\n", kfsession)
}

func TestWechatMp_GetKfsessionList(t *testing.T) {
	kfsessionList, _ := customer_service_wechatMp.GetKfsessionList(customer_service_token, "openid")
	fmt.Printf("%+v\n", kfsessionList)
}

func TestWechatMp_GetWaitcaseList(t *testing.T) {
	waitcaseList, _ := customer_service_wechatMp.GetWaitcaseList(customer_service_token)
	fmt.Printf("%+v\n", waitcaseList)
}

func TestWechatMp_GetMsgrecordList(t *testing.T) {
	startTime := time.Now().Unix()
	after, _ := time.ParseDuration("1h")
	endTime := time.Now().Add(after).Unix()
	var param = &GetMsgListParam{
		Starttime: startTime,
		Endtime:   endTime,
		Msgid:     1,
		Number:    10000,
	}
	getMsgListResp, _ := customer_service_wechatMp.GetMsgrecordList(customer_service_token, param)
	fmt.Printf("%+v\n", getMsgListResp)
}
