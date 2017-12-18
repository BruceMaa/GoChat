package mp

import (
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"github.com/chanxuehong/wechat/json"
)

const (
	WECHAT_USER_TAGS_CREATE_API                 = `https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s`                   // 创建标签
	WECHAT_USER_TAGS_GET_API                    = `https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s`                      // 获取公众号已创建的标签
	WECHAT_USER_TAGS_UPDATE_API                 = `https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s`                   // 编辑标签
	WECHAT_USER_TAGS_DELETE_API                 = `https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s`                   // 删除标签
	WECHAT_USER_TAG_USERS_API                   = `https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s`                  // 获取标签下粉丝列表
	WECHAT_USER_TAGS_MEMBERS_BATCHTAGGING_API   = `https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s`     // 批量为用户打标签
	WECHAT_USER_TAGS_MEMBERS_BATCHUNTAGGING_API = `https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s`   // 批量为用户取消标签
	WECHAT_USER_TAGS_MEMBERS_BLACKLIST_API      = `https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s`     // 获取公众号的黑名单列表
	WECHAT_USER_TAGS_MEMBERS_BATCHBLACK_API     = `https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s`   // 批量拉黑用户
	WECHAT_USER_TAGS_MEMBERS_BATCHUNBLACK_API   = `https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=%s` // 批量取消拉黑用户
	WECHAT_USER_TAGS_GETIDLIST_API              = `https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s`                // 获取用户身上的标签列表

	WECHAT_USER_UPDATE_REMARK_API = `https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%s` // 设置用户备注名

	WECHAT_USER_INFO_API       = `https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s` // 获取用户基本信息
	WECHAT_USER_INFO_BATCH_API = `https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s`          // 批量获取用户基本信息

	WECHAT_USER_GET_API = `https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s` // 获取用户列表
)

type WechatSex int

const (
	WECHAT_SEX_UNKNOW WechatSex = iota // 未知性别
	WECHAT_SEX_MALE                    // 男性
	WECHAT_SEX_FEMALE                  // 女性
)

type (
	WechatUserInfo struct {
		Subscribe     int       `json:"subscribe"`      // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
		Openid        string    `json:"openid"`         // 用户的标识，对当前公众号唯一
		Nickname      string    `json:"nickname"`       // 用户的昵称
		Sex           WechatSex `json:"sex"`            // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
		City          string    `json:"city"`           // 用户所在城市
		Country       string    `json:"country"`        // 用户所在国家
		Province      string    `json:"province"`       // 用户所在省份
		Language      string    `json:"language"`       // 用户的语言，简体中文为zh_CN
		Headimgurl    string    `json:"headimgurl"`     // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
		SubscribeTime int64     `json:"subscribe_time"` // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
		Unionid       string    `json:"unionid"`        // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
		Remark        string    `json:"remark"`         // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
		Groupid       int       `json:"groupid"`        // 用户所在的分组ID（兼容旧的用户分组接口）
		TagidList     []int     `json:"tagid_list"`     // 用户被打上的标签ID列表
		common.PublicError
	}

	WechatUserList struct {
		UserList []WechatUserInfo `json:"user_list"`
	}

	WechatUserInfoList struct {
		UserInfoList []WechatUserInfo `json:"user_info_list"`
		common.PublicError
	}

	WechatUserOpenIdList struct {
		Total int `json:"total"` // 关注该公众账号的总用户数
		Count int `json:"count"` // 拉取的OPENID个数，最大值为10000
		Data  struct {
			Openid []string `json:"openid"`
		} `json:"data"` // 列表数据，OPENID的列表
		NextOpenid string `json:"next_openid"` // 拉取列表的最后一个用户的OPENID
		common.PublicError
	}

	WechatUserTags struct {
		Tags []WechatUserTag `json:"tags"`
	}
	WechatUserTag struct {
		Id    int    `json:"id,omitempty"`    // 标签id，由微信分配
		Name  string `json:"name"`            // 标签名，UTF8编码 （30个字符以内）
		Count int    `json:"count,omitempty"` // 此标签下粉丝数
		common.PublicError
	}
)

// 获取用户基本信息
func (wm *WechatMp) UserInfo(accessToken, openid string) (*WechatUserInfo, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_USER_INFO_API, accessToken, openid, WECHAT_LANGUAGE_ZH_CN))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UserInfo http get error: %+v\n", err)
		return nil, err
	}
	var userInfo WechatUserInfo
	if err = json.Unmarshal(resp, &userInfo); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UserInfo json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &userInfo, err
}

// 批量获取用户基本信息
func (wm *WechatMp) BatchUserInfos(accessToken string, userList *WechatUserList) (*WechatUserInfoList, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_USER_INFO_BATCH_API, accessToken), userList)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchUserInfos http post error: %+v\n", err)
		return nil, err
	}
	var userInfoList WechatUserInfoList
	if err = json.Unmarshal(resp, &userInfoList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchUserInfos json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &userInfoList, nil
}

// 获取用户列表
func (wm *WechatMp) GetUsers(accessToken, nextOpenId string) (*WechatUserOpenIdList, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_USER_GET_API, accessToken, nextOpenId))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Get Users http get error: %+v\n", err)
		return nil, err
	}
	var wechatUserOpenIdList WechatUserOpenIdList
	if err := json.Unmarshal(resp, &wechatUserOpenIdList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Get Users json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &wechatUserOpenIdList, nil
}

// 创建标签
func (wm *WechatMp) CreateTag(accessToken string, tag *WechatUserTag) (*WechatUserTag, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_USER_TAGS_CREATE_API, accessToken), tag)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Create Tag http post error: %+v\n", err)
		return nil, err
	}
	var resultTag WechatUserTag
	if err = json.Unmarshal(resp, &resultTag); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Create Tag json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &resultTag, nil
}

// 获取公众号已创建的标签
func (wm *WechatMp) GetTags(accessToken string) (*WechatUserTags, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_USER_TAGS_GET_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Get Tags http get error: %+v\n", err)
		return nil, err
	}
	var userTags WechatUserTags
	if err = json.Unmarshal(resp, &userTags); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Get Tags json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &userTags, nil
}

// 编辑标签
func (wm *WechatMp) UpdateTag(accessToken string) {
	//TODO
}

// 删除标签
func (wm *WechatMp) DeleteTag(accessToken string) {
	//TODO
}

// 获取标签下粉丝列表
func (wm *WechatMp) GetUsersWithTag(accessToken string) {
	//TODO
}

// 批量为用户打标签
func (wm *WechatMp) BatchTagging(accessToken string) {
	//TODO
}

// 批量为用户取消标签
func (wm *WechatMp) BatchUnTagging(accessToken string) {
	//TODO
}

// 获取用户身上的标签列表
func (wm *WechatMp) GetTagsWithUser(accessToken string) {
	//TODO
}

// 设置用户备注名
func (wm *WechatMp) UpdateUserRemark(accessToken string) {
	//TODO
}

// 获取公众号的黑名单列表
func (wm *WechatMp) GetBlackList(accessToken, beginOpenId string) (*WechatUserOpenIdList, error) {
	type blackListReq struct {
		BeginOpenid string `json:"begin_openid"` // 当 begin_openid 为空时，默认从开头拉取。
	}
	blackReq := &blackListReq{
		BeginOpenid: beginOpenId,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_USER_TAGS_MEMBERS_BLACKLIST_API, accessToken), *blackReq)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetBlackList http post error: %+v\n", err)
		return nil, err
	}
	var wechatUserOpenIdList WechatUserOpenIdList
	if err = json.Unmarshal(resp, &wechatUserOpenIdList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetBlackList json.Unmarshal error: %+v\n", err)
		return nil, err
	}

	return &wechatUserOpenIdList, nil
}

type openidListReq struct {
	OpenidList []string `json:"openid_list"`
}

// 批量拉黑用户
func (wm *WechatMp) BatchBlackUsers(accessToken string, openidList ...string) (*common.PublicError, error) {
	openidReq := &openidListReq{
		OpenidList: openidList,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_USER_TAGS_MEMBERS_BATCHBLACK_API, accessToken), &openidReq)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchBlackUsers http post error: %+v\n", err)
		return nil, err
	}
	var publicError common.PublicError
	if err = json.Unmarshal(resp, &publicError); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchBlackUsers json.Unmarshal error: %+v\n", err)
		return nil, err
	}

	return &publicError, nil
}

// 批量取消拉黑用户
func (wm *WechatMp) BatchUnBlackUsers(accessToken string, openidList ...string) (*common.PublicError, error) {
	openidReq := &openidListReq{
		OpenidList: openidList,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_USER_TAGS_MEMBERS_BATCHUNBLACK_API, accessToken), &openidReq)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchUnBlackUsers http post error: %+v\n", err)
		return nil, err
	}
	var publicError common.PublicError
	if err = json.Unmarshal(resp, &publicError); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchUnBlackUsers json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &publicError, nil
}
