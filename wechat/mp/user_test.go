package mp

import (
	"fmt"
	"testing"
)

var user_tags_token string
var user_tags_wechatMp WechatMp

func init() {
	config := &WechatMpConfig{
		AppId:     "wx5fa42349ef54acfc",
		AppSecret: "4f1c8ee9007b9aa71bca7a542e659483",
		Token:     "bingobox",
	}
	user_tags_wechatMp.Configure = *config
	accessToken, err := menu_wechatMp.AccessTokenFromWechat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accessToken)
	user_tags_token = accessToken.AccessToken
}

func TestWechatMp_UserInfo(t *testing.T) {
	userInfo, _ := user_tags_wechatMp.UserInfo(user_tags_token, "ohMaywPoQB6h7mhRtuvMBAZ49Wso")
	fmt.Printf("userInfo: %+v\n", userInfo)
}

func TestWechatMp_GetUsers(t *testing.T) {
	var token = `4_GJ6q40cZzNX2AdQlEvMk0RXmX6pDCtSgzIo8h00NKRiep5bzPUexhabhVU-V2PuJheMa_Lr3Dk_-RFRUU6ZPqx1iUFq7dH7HunhSs3eELvu1t7N6kOuT8CcXFIQr8CUV6f50W3xZ3afJBav5EDVcAFAMLD`
	openidList, _ := user_tags_wechatMp.GetUsers(token, "oFBlPv6P9UuS8mliE2ZPqx-Di0M4")
	fmt.Printf("%+v\n", openidList)
}

func TestWechatMp_BatchUserInfos(t *testing.T) {
	userInfo1 := &WechatUserInfo{
		Openid: "ohMaywPoQB6h7mhRtuvMBAZ49Wso",
	}
	var userList = []WechatUserInfo{
		*userInfo1,
	}
	wechatUserList := &WechatUserList{
		UserList: userList,
	}
	userInfoList, _ := user_tags_wechatMp.BatchUserInfos(user_tags_token, wechatUserList)
	fmt.Printf("userInfoList: %+v\n", userInfoList)
}

func TestTagsCreate(t *testing.T) {
	wechatUserTag := &WechatUserTag{}
	wechatUserTag.Name = "标签3"
	tag, _ := user_tags_wechatMp.CreateTag(user_tags_token, wechatUserTag)
	fmt.Printf("wechatUserTag: %+v\n", tag)
}

func TestTagsGet(t *testing.T) {
	tags, _ := user_tags_wechatMp.GetTags(user_tags_token)
	fmt.Printf("wechatUserTags: %+v\n", tags)
}
