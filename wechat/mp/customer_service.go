package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
)

const (
	GETKFLIST_CUSTOMERSERVICE_API               = `https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s`                    // 获取所有客服账号
	GETONLINEKFLIST_CUSTOMERSERVICE_API         = `https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=%s`              // 获取在线客服账号
	ADD_KFACCOUNT_CUSTOMERSERVICE_API           = `https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s`                        // 添加客服帐号
	INVITEWORKER_KFACCOUNT_CUSTOMERSERVICE_API  = `https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=%s`               // 邀请绑定客服帐号
	UPDATE_KFACCOUNT_CUSTOMERSERVICE_API        = `https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s`                     // 设置客服信息
	UPLOADHEADIMG_KFACCOUNT_CUSTOMERSERVICE_API = `http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s` // 上传客服头像
	DEL_KFACCOUNT_CUSTOMERSERVICE_API           = `https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s`                        // 删除客服帐号

	CREATE_KFSESSION_CUSTOMERSERVICE_API       = `https://api.weixin.qq.com/customservice/kfsession/create?access_token=%s`                       // 创建会话
	CLOSE_KFSESSION_CUSTOMERSERVICE_API        = `https://api.weixin.qq.com/customservice/kfsession/close?access_token=%s`                        // 关闭会话
	GET_KFSESSION_CUSTOMERSERVICE_API          = `https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=%s&openid=%s`         // 获取客户会话状态
	GET_KFSESSIONLIST_CUSTOMERSERVICE_API      = `https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=%s&kf_account=%s` // 获取客服会话列表
	GET_WAIT_KFSESSIONLIST_CUSTOMERSERVICE_API = `https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=%s`                  // 获取未接入会话列表
	GET_MSGLIST_MSGRECORD_CUSTOMERSERVICE_API  = `https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=%s`                   // 获取聊天记录

	SEND_CUSTOM_MESSAGE_API = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s` // 客服接口-发消息
)

type (
	CustomerServiceInfoList struct {
		KfList []CustomerServiceInfo `json:"kf_list,omitempty"`
		common.PublicError
	}
	CustomerServiceInfo struct {
		KfAccount        string `json:"kf_account"`                   // 完整客服帐号，格式为：帐号前缀@公众号微信号
		KfNick           string `json:"kf_nick"`                      // 客服昵称
		KfId             string `json:"kf_id"`                        // 客服编号
		KfHeadimgurl     string `json:"kf_headimgurl"`                // 客服头像
		KfWx             string `json:"kf_wx,omitempty"`              // 如果客服帐号已绑定了客服人员微信号，则此处显示微信号
		InviteWx         string `json:"invite_wx,omitempty"`          // 如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请，则此处显示绑定邀请的微信号
		InviteExpireTime string `json:"invite_expire_time,omitempty"` // 如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请，邀请的过期时间，为unix 时间戳
		InviteStatus     string `json:"invite_status,omitempty"`      // 邀请的状态，有等待确认“waiting”，被拒绝“rejected”，过期“expired”
	}

	CustomerServiceInfoOnlineList struct {
		KfOnlineList []CustomerServiceInfoOnline `json:"kf_online_list"`
		common.PublicError
	}
	CustomerServiceInfoOnline struct {
		KfAccount    string `json:"kf_account"`    // 完整客服帐号，格式为：帐号前缀@公众号微信号
		Status       int    `json:"status"`        // 客服在线状态，目前为：1、web 在线
		KfId         string `json:"kf_id"`         // 客服编号
		AcceptedCase int    `json:"accepted_case"` // 客服当前正在接待的会话数
	}

	KfInfo struct {
		KfAccount string `json:"kf_account"`          // 完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
		Nickname  string `json:"nickname,omitempty"`  // 客服昵称，最长16个字
		Headimg   string `json:"headimg"`             // 客服头像
		InviteWx  string `json:"invite_wx,omitempty"` // 接收绑定邀请的客服微信号
	}

	Kfsession struct {
		KfAccount  string `json:"kf_account"`           // 完整客服帐号，格式为：帐号前缀@公众号微信号
		Openid     string `json:"openid"`               // 粉丝的openid
		Createtime int64  `json:"createtime,omitempty"` // 会话接入的时间
		LatestTime int64  `json:"latest_time"`          // 粉丝的最后一条消息的时间
		common.PublicError
	}

	KfsessionList struct {
		Sessionlist []Kfsession `json:"sessionlist"`
		common.PublicError
	}

	WaitcaseList struct {
		Count        int         `json:"count"`        // 未接入会话数量
		Waitcaselist []Kfsession `json:"waitcaselist"` // 未接入会话列表，最多返回100条数据，按照来访顺序
		common.PublicError
	}

	GetMsgListParam struct {
		Starttime int64 `json:"starttime"` // 起始时间，unix时间戳
		Endtime   int64 `json:"endtime"`   // 结束时间，unix时间戳，每次查询时段不能超过24小时
		Msgid     int64 `json:"msgid"`     // 消息id顺序从小到大，从1开始
		Number    int64 `json:"number"`    // 每次获取条数，最多10000条
	}

	GetMsgListResp struct {
		Worker   string `json:"worker"`   // 完整客服帐号，格式为：帐号前缀@公众号微信号
		Openid   string `json:"openid"`   // 用户标识
		Opercode int    `json:"opercode"` // 操作码，2002（客服发送信息），2003（客服接收消息）
		Text     string `json:"text"`     // 聊天记录
		Time     int64  `json:"time"`     // 操作时间，unix时间戳
		common.PublicError
	}
)

// 客服消息类型
const (
	CUSTOMER_MSG_TYPE_TEXT  = "text"
	CUSTOMER_MSG_TYPE_IMAGE = "image"
	CUSTOMER_MSG_TYPE_VOICE = "voice"
	CUSTOMER_MSG_TYPE_VIDEO = "video"
)

type (
	CustomerMessageRequest struct {
		Touser          string `json:"touser"` // 普通用户openid
		Msgtype         string `json:"msgtype"`
		Customerservice struct {
			KfAccount string `json:"kf_account"`
		} `json:"customerservice,omitempty"` // 客服账号
	}

	CustomerTextMessageRequest struct {
		CustomerMessageRequest
		Text struct {
			Content string `json:"content"` // 文本消息
		} `json:"text"`
	}

	CustomerImageMessageRequest struct {
		CustomerMessageRequest
		Image struct {
			MediaId string `json:"media_id"` // 图片消息
		} `json:"image"`
	}

	CustomerVoiceMessageRequest struct {
		CustomerMessageRequest
		Voice struct {
			MediaId string `json:"media_id"` // 语音消息
		} `json:"voice"`
	}

	CustomerVideoMessageRequest struct {
		CustomerMessageRequest
		Video struct {
			MediaId      string `json:"media_id"`
			ThumbMediaId string `json:"thumb_media_id"`
			Title        string `json:"title"`
			Description  string `json:"description"`
		} `json:"video"`
	}
)

// 获取客服基本信息
func (wm *WechatMp) GetKfList(accessToken string) (*CustomerServiceInfoList, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(GETKFLIST_CUSTOMERSERVICE_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "getkflist error: %+v", err)
		return nil, fmt.Errorf("getkflist error: %+v", err)
	}
	var customerServiceInfoList CustomerServiceInfoList
	if err = json.Unmarshal([]byte(resp), &customerServiceInfoList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "getkflist error: %+v", err)
		return nil, fmt.Errorf("getkflist error: %+v", err)
	}
	return &customerServiceInfoList, nil
}

// 获取客服基本信息
func (wm *WechatMp) GetOnlineKfList(accessToken string) (*CustomerServiceInfoOnlineList, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(GETONLINEKFLIST_CUSTOMERSERVICE_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "getonlinekflist error: %+v", err)
		return nil, fmt.Errorf("getonlinekflist error: %+v", err)
	}
	var customerServiceInfoOnlineList CustomerServiceInfoOnlineList
	if err = json.Unmarshal([]byte(resp), &customerServiceInfoOnlineList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "getonlinekflist error: %+v", err)
		return nil, fmt.Errorf("getonlinekflist error: %+v", err)
	}
	return &customerServiceInfoOnlineList, nil
}

// 添加客服账号
func (wm *WechatMp) AddKfaccount(accessToken string, kfInfo *KfInfo) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(ADD_KFACCOUNT_CUSTOMERSERVICE_API, accessToken), &kfInfo)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "add kfaccount error: %+v", err)
		return nil, fmt.Errorf("add kfaccount error: %+v", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "add kfaccount error: %+v", err)
		return nil, fmt.Errorf("add kfaccount error: %+v", err)
	}
	return &result, nil
}

// 邀请绑定客服帐号
func (wm *WechatMp) InviteKfaccount(accessToken string, kfInfo *KfInfo) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(INVITEWORKER_KFACCOUNT_CUSTOMERSERVICE_API, accessToken), &kfInfo)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "invite kfaccount error: %+v", err)
		return nil, fmt.Errorf("invite kfaccount error: %+v", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "invite kfaccount error: %+v", err)
		return nil, fmt.Errorf("invite kfaccount error: %+v", err)
	}
	return &result, nil
}

// 设置客服信息
func (wm *WechatMp) UpdateKfaccount(accessToken string, kfInfo *KfInfo) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(UPDATE_KFACCOUNT_CUSTOMERSERVICE_API, accessToken), &kfInfo)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "update kfaccount error: %+v", err)
		return nil, fmt.Errorf("update kfaccount error: %+v", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "update kfaccount error: %+v", err)
		return nil, fmt.Errorf("update kfaccount error: %+v", err)
	}
	return &result, nil
}

// TODO 上传客服头像
func (wm *WechatMp) UploadheadimgKfaccount(accessToken string, kfInfo *KfInfo) (*common.PublicError, error) {
	//common.HTTPPostForm()
	return nil, nil
}

// 删除客服账号
func (wm *WechatMp) DeleteKfaccount(accessToken string, kfInfo *KfInfo) (*common.PublicError, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(DEL_KFACCOUNT_CUSTOMERSERVICE_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "delete kfaccount error: %+v\n", err)
		return nil, fmt.Errorf("delete kfaccount error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "delete kfaccount error: %+v\n", err)
		return nil, fmt.Errorf("delete kfaccount error: %+v\n", err)
	}
	return &result, nil
}

// 创建会话
func (wm *WechatMp) CreateKfsession(accessToken string, kfsession *Kfsession) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(CREATE_KFSESSION_CUSTOMERSERVICE_API, accessToken), &kfsession)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "create kfsession error: %+v\n", err)
		return nil, fmt.Errorf("create kfsession error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "create kfsession error: %+v\n", err)
		return nil, fmt.Errorf("create kfsession error: %+v\n", err)
	}
	return &result, nil
}

// 关闭会话
func (wm *WechatMp) CloseKfsession(accessToken string, kfsession *Kfsession) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(CLOSE_KFSESSION_CUSTOMERSERVICE_API, accessToken), &kfsession)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "close kfsession error: %+v\n", err)
		return nil, fmt.Errorf("close kfsession error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "close kfsession error: %+v\n", err)
		return nil, fmt.Errorf("close kfsession error: %+v\n", err)
	}
	return &result, nil
}

// 获取客户会话状态
func (wm *WechatMp) GetKfsession(accessToken, openid string) (*Kfsession, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(GET_KFSESSION_CUSTOMERSERVICE_API, accessToken, openid))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get kfsession error: %+v\n", err)
		return nil, fmt.Errorf("get kfsession error: %+v\n", err)
	}
	var kfsession Kfsession
	if err = json.Unmarshal([]byte(resp), &kfsession); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get kfsession error: %+v\n", err)
		return nil, fmt.Errorf("get kfsession error: %+v\n", err)
	}
	return &kfsession, nil
}

// 获取客服会话列表
func (wm *WechatMp) GetKfsessionList(accessToken, kfAccount string) (*KfsessionList, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(GET_KFSESSIONLIST_CUSTOMERSERVICE_API, accessToken, kfAccount))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get kfsession list error: %+v\n", err)
		return nil, fmt.Errorf("get kfsession list error: %+v\n", err)
	}
	var kfsessionList KfsessionList
	if err = json.Unmarshal([]byte(resp), &kfsessionList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get kfsession list error: %+v\n", err)
		return nil, fmt.Errorf("get kfsession list error: %+v\n", err)
	}
	return &kfsessionList, nil
}

// 获取未接入会话列表
func (wm *WechatMp) GetWaitcaseList(accessToken string) (*WaitcaseList, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(GET_WAIT_KFSESSIONLIST_CUSTOMERSERVICE_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get waitcase list error: %+v", err)
		return nil, fmt.Errorf("get waitcase list error: %+v", err)
	}
	var waitcaseList WaitcaseList
	if err = json.Unmarshal([]byte(resp), &waitcaseList); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get waitcase list error: %+v", err)
		return nil, fmt.Errorf("get waitcase list error: %+v", err)
	}
	return &waitcaseList, nil
}

// 获取聊天记录
func (wm *WechatMp) GetMsgrecordList(accessToken string, param *GetMsgListParam) (*GetMsgListResp, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(GET_MSGLIST_MSGRECORD_CUSTOMERSERVICE_API, accessToken), &param)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get msg record list error: %+v\n", err)
		return nil, fmt.Errorf("get msg record list error: %+v\n", err)
	}
	var getMsgListResp GetMsgListResp
	if err = json.Unmarshal([]byte(resp), &getMsgListResp); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "get msg record list error: %+v\n", err)
		return nil, fmt.Errorf("get msg record list error: %+v\n", err)
	}
	return &getMsgListResp, nil
}

// ********************************************发送消息********************************

// 发送客服消息
func (wm *WechatMp) SendCustomerMessage(accessToken string, message interface{}) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(SEND_CUSTOM_MESSAGE_API, accessToken), &message)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "send custom message error: %+v\n", err)
		return nil, fmt.Errorf("send custom message error: %+v\n", err)
	}
	var result common.PublicError
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "send custom message error: %+v\n", err)
		return nil, fmt.Errorf("send custom message error: %+v\n", err)
	}
	return &result, nil
}

// 客服发送文本消息
func (wm *WechatMp) SendCustomerTextMessage(accessToken string, textMessage *CustomerTextMessageRequest) (*common.PublicError, error) {
	textMessage.Msgtype = CUSTOMER_MSG_TYPE_TEXT
	return wm.SendCustomerMessage(accessToken, *textMessage)
}

// 客服发送图片消息
func (wm *WechatMp) SendCustomerImageMessage(accessToken string, imageMessage *CustomerImageMessageRequest) (*common.PublicError, error) {
	imageMessage.Msgtype = CUSTOMER_MSG_TYPE_IMAGE
	return wm.SendCustomerMessage(accessToken, *imageMessage)
}

// 客服发送语音消息
func (wm *WechatMp) SendCustomerVoiceMessage(accessToken string, voiceMessage *CustomerVoiceMessageRequest) (*common.PublicError, error) {
	voiceMessage.Msgtype = CUSTOMER_MSG_TYPE_VOICE
	return wm.SendCustomerMessage(accessToken, *voiceMessage)
}

// 客服发送视频消息
func (wm *WechatMp) SendCustomerVideoMessage(accessToken string, videoMessage *CustomerVideoMessageRequest) (*common.PublicError, error) {
	videoMessage.Msgtype = CUSTOMER_MSG_TYPE_VIDEO
	return wm.SendCustomerMessage(accessToken, *videoMessage)
}
