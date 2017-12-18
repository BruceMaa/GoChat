package mp

import (
	"fmt"
	"testing"
)

var message_token string
var message_wechatMp WechatMp

func init() {
	config := &WechatMpConfig{
		AppId:     "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token:     "bingobox",
	}
	message_wechatMp.Configure = *config
	accessToken, err := message_wechatMp.AccessTokenFromWechat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
	message_token = accessToken.AccessToken
}

func TestWechatMp_SetTemplateIndustries(t *testing.T) {
	var industries = [2]int{2}
	message_wechatMp.SetTemplateIndustries(message_token, industries)
}

func TestWechatMp_GetTemplateIndustries(t *testing.T) {
	templateIndustries, _ := message_wechatMp.GetTemplateIndustries(message_token)
	fmt.Printf("%+v\n", templateIndustries)
}

func TestWechatMp_GetAllTemplates(t *testing.T) {
	templates, _ := message_wechatMp.GetAllTemplates(message_token)
	fmt.Printf("%+v\n", templates)
}

func TestWechatMp_AddTemplate(t *testing.T) {
	//message_wechatMp.AddTemplate()
}

func TestWechatMp_SendTemplate(t *testing.T) {
	sendTemplate := &SendTemplate{
		Touser:     "ohMaywPoQB6h7mhRtuvMBAZ49Wso",
		TemplateId: "HGEeYrY3ysxgU60a0KApiC6sob0vba16SCgSseXFs7I",
		Url:        "www.qq.com",
	}
	data := make(map[string]SendTemplateData)
	sendTemplateData1 := &SendTemplateData{
		Value: "马强",
		Color: "#173177",
	}
	data["name"] = *sendTemplateData1
	sendTemplateData2 := &SendTemplateData{
		Value: "maqiang@mail.com",
		Color: "#173177",
	}
	data["email"] = *sendTemplateData2
	sendTemplateData3 := &SendTemplateData{
		Value: "18",
	}
	data["age"] = *sendTemplateData3
	sendTemplateData4 := &SendTemplateData{
		Value: "备注",
	}
	data["remark"] = *sendTemplateData4
	sendTemplate.Data = &data
	resp, _ := message_wechatMp.SendTemplate(message_token, sendTemplate)
	fmt.Printf("resp: %+v\n", *resp)
}

func TestWechatMp_DeleteTemplate(t *testing.T) {
	resp, _ := message_wechatMp.DeleteTemplate(message_token, "HGEeYrY3ysxgU60a0KApiC6sob0vba16SCgSseXFs7I")
	fmt.Printf("resp: %+v\n", resp)
}

func TestWechatMp_BuildSubscribeMsgURL(t *testing.T) {
	resp := message_wechatMp.BuildSubscribeMsgURL(12, "zQAaB1NoLmAkNf0enpBAlh74XdCYmZfkvEwUObTNpiM", "http://wechat.bingobox.cc/wechat/subscribe", "mark")
	fmt.Printf("resp: %s\n", resp)
}
