package mp

import (
	"encoding/xml"
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"github.com/chanxuehong/wechat/json"
	"time"
)

const (
	MSG_TYPE_TEXT        = "text"       //文本消息
	MSG_TYPE_IMG         = "image"      //图片消息
	MSG_TYPE_VOICE       = "voice"      //声音消息
	MSG_TYPE_VIDEO       = "video"      //视频消息
	MSG_TYPE_VIDEO_SHORT = "shortvideo" //短视频消息
	MSG_TYPE_LOCATION    = "location"   //地理位置消息
	MSG_TYPE_LINK        = "link"       //链接消息
	MSG_TYPE_MUSIC       = "music"      //音乐消息
	MSG_TYPE_NEWS        = "news"       //图文消息
	MSG_TYPE_EVENT       = "event"      //事件消息
)

const (
	MSG_EVENT_SUBSCRIBE                    = "subscribe"                    // 关注事件
	MSG_EVENT_UNSUBSCRIBE                  = "unsubscribe"                  // 取消关注事件
	MSG_EVENT_SCAN                         = "SCAN"                         // 扫码二维码事件
	MSG_EVENT_LOCATION                     = "LOCATION"                     // 上报地理位置事件
	MSG_EVENT_CLICK                        = "CLICK"                        // 点击自定义菜单拉取消息事件
	MSG_EVENT_VIEW                         = "VIEW"                         // 点击自定义菜单跳转链接事件
	MSG_EVENT_QUALIFICATION_VERIFY_SUCCESS = "qualification_verify_success" // 资质认证成功通知
	MSG_EVENT_QUALIFICATION_VERIFY_FAIL    = "qualification_verify_fail"    // 资质认证失败通知
	MSG_EVENT_NAMING_VERIFY_SUCCESS        = "naming_verify_success"        // 名称认证成功
	MSG_EVENT_NAMING_VERIFY_FAIL           = "naming_verify_fail"           // 名称认证失败
	MSG_EVENT_ANNUAL_RENEW                 = "annual_renew"                 // 年审通知
	MSG_EVENT_VERIFY_EXPIRED               = "verify_expired"               // 认证过期失效通知
	MSG_EVENT_TEMPLATESENDJOBFINISH        = "TEMPLATESENDJOBFINISH"        // 发送模板消息结果通知
)

type (
	MsgEventSubscribe struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID）
		CreateTime   int64    // 消息创建时间 （整型）
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
		EventKey     string   // 事件KEY值，如果有qrscene_为前缀，则表示为扫描二维码关注
		Ticket       string   // 二维码的ticket，可用来换取二维码图片
	}

	MsgEventScan struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID）
		CreateTime   int64    // 消息创建时间 （整型）
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型，SCAN
		EventKey     string   // 事件KEY值，即创建二维码时的二维码scene_id，或者scene_str
		Ticket       string   // 二维码的ticket，可用来换取二维码图片
	}

	MsgEventLocation struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID）
		CreateTime   int64    // 消息创建时间 （整型）
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型，LOCATION
		Latitude     string   // 地理位置纬度
		Longitude    string   // 地理位置经度
		Precision    string   // 地理位置精度
	}

	MsgEventMenuClick struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID）
		CreateTime   int64    // 消息创建时间 （整型）
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型，CLICK
		EventKey     string   // 事件KEY值，与自定义菜单接口中KEY值对应
	}

	MsgEventMenuView struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID）
		CreateTime   int64    // 消息创建时间 （整型）
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型，VIEW
		EventKey     string   // 事件KEY值，设置的跳转URL
	}

	MsgEventQualificationVerifySuccess struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID，此时发送方是系统帐号）
		CreateTime   int64    // 消息创建时间 （整型），时间戳
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 qualification_verify_success
		ExpiredTime  int64    // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
	}

	MsgEventQualificationVerifyFail struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID，此时发送方是系统帐号）
		CreateTime   int64    // 消息创建时间 （整型），时间戳
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 qualification_verify_fail
		FailTime     string   // 失败发生时间 (整形)，时间戳
		FailReason   string   // 认证失败的原因
	}

	MsgEventNamingVerifySuccess struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID，此时发送方是系统帐号）
		CreateTime   int64    // 消息创建时间 （整型），时间戳
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 naming_verify_success
		ExpiredTime  int64    // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
	}

	MsgEventNamingVerifyFail struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID，此时发送方是系统帐号）
		CreateTime   int64    // 消息创建时间 （整型），时间戳
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 naming_verify_fail
		FailTime     string   // 失败发生时间 (整形)，时间戳
		FailReason   string   // 认证失败的原因
	}

	MsgEventAnnualRenew struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID，此时发送方是系统帐号）
		CreateTime   int64    // 消息创建时间 （整型），时间戳
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 annual_renew，提醒公众号需要去年审了
		ExpiredTime  int64    // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
	}

	MsgEventVerifyExpired struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 开发者微信号
		FromUserName string   // 发送方帐号（一个OpenID，此时发送方是系统帐号）
		CreateTime   int64    // 消息创建时间 （整型），时间戳
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 verify_expired
		ExpiredTime  int64    // 有效期 (整形)，指的是时间戳，表示已于该时间戳认证过期，需要重新发起微信认证
	}

	MsgEventTemplateSendFinish struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string   // 公众号微信号
		FromUserName string   // 接收模板消息的用户的openid
		CreateTime   int64    // 创建时间
		MsgType      string   // 消息类型，event
		Event        string   // 事件类型 TEMPLATESENDJOBFINISH
		MsgID        int64    // 消息ID
		Status       string   // 发送状态为成功:success, 发送状态为用户拒绝接收:failed:user block, 发送状态为发送失败（非用户拒绝）:failed: system failed
	}
)

// 防止回复消息时，出现xml格式错误
type CDATAText struct {
	Text string `xml:",innerxml"`
}

type (
	MsgTextResponse struct {
		XMLName      xml.Name  `xml:"xml"`
		ToUserName   CDATAText // 接收方帐号（收到的OpenID）
		FromUserName CDATAText // 开发者微信号
		CreateTime   int64     // 消息创建时间 （整型）
		MsgType      CDATAText // text
		Content      CDATAText // 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
	}

	MsgImageResponse struct {
		XMLName      xml.Name  `xml:"xml"`
		ToUserName   CDATAText // 接收方帐号（收到的OpenID）
		FromUserName CDATAText // 开发者微信号
		CreateTime   int64     // 消息创建时间 （整型）
		MsgType      CDATAText // image
		Image        struct {
			MediaId CDATAText // 通过素材管理中的接口上传多媒体文件，得到的id。
		}
	}

	MsgVoiceResponse struct {
		XMLName      xml.Name  `xml:"xml"`
		ToUserName   CDATAText // 接收方帐号（收到的OpenID）
		FromUserName CDATAText // 开发者微信号
		CreateTime   int64     // 消息创建时间戳 （整型）
		MsgType      CDATAText // 语音，voice
		Voice        struct {
			MediaId CDATAText // 通过素材管理中的接口上传多媒体文件，得到的id
		}
	}

	MsgVideoResponse struct {
		XMLName      xml.Name  `xml:"xml"`
		ToUserName   CDATAText // 接收方帐号（收到的OpenID）
		FromUserName CDATAText // 开发者微信号
		CreateTime   int64     // 消息创建时间 （整型）
		MsgType      CDATAText // video
		Video        struct {
			MediaId     CDATAText // 通过素材管理中的接口上传多媒体文件，得到的id
			Title       CDATAText // 视频消息的标题
			Description CDATAText // 视频消息的描述
		}
	}

	MsgMusicResponse struct {
		XMLName      xml.Name  `xml:"xml"`
		ToUserName   CDATAText // 接收方帐号（收到的OpenID）
		FromUserName CDATAText // 开发者微信号
		CreateTime   int64     // 消息创建时间 （整型）
		MsgType      CDATAText // music
		Music        struct {
			Title        CDATAText // 音乐标题
			Description  CDATAText // 音乐描述
			MusicUrl     CDATAText // 音乐链接
			HQMusicUrl   CDATAText // 高质量音乐链接，WIFI环境优先使用该链接播放音乐
			ThumbMediaId CDATAText // 缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id
		}
	}

	MsgNewsResponse struct {
		XMLName      xml.Name              `xml:"xml"`
		ToUserName   CDATAText             // 接收方帐号（收到的OpenID）
		FromUserName CDATAText             // 开发者微信号
		CreateTime   int64                 // 消息创建时间 （整型）
		MsgType      CDATAText             // news
		ArticleCount int                   // 图文消息个数，限制为8条以内
		Articles     []MsgNewsItemResponse // 多条图文消息信息，默认第一个item为大图,注意，如果图文数超过8，则将会无响应
	}

	MsgNewsItemResponse struct {
		XMLName     xml.Name  `xml:"item"`
		Title       CDATAText // 图文消息标题
		Description CDATAText // 图文消息描述
		PicUrl      CDATAText // 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
		Url         CDATAText // 点击图文消息跳转链接
	}
)

type (
	TextMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // text
		Content      string // 文本消息内容
		MsgId        int64  // 消息id，64位整型
	}

	ImageMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // image
		PicUrl       string // 图片链接（由系统生成）
		MediaId      string // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
		MsgId        int64  // 消息id，64位整型
	}

	VoiceMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // voice
		MediaId      string // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
		Format       string // 语音格式，如amr，speex等
		MsgID        int64  // 消息id，64位整型
	}

	VideoMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // video
		MediaId      string // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
		ThumbMediaId string // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
		MsgId        int64  // 消息id，64位整型
	}

	ShortVideoMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // shortvideo
		MediaId      string // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
		ThumbMediaId string // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
		MsgId        int64  // 消息id，64位整型
	}

	LocationMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // shortvideo
		Location_X   string // 地理位置维度
		Location_Y   string // 地理位置经度
		Scale        int    // 地图缩放大小
		Label        string // 地理位置信息
		MsgId        int64  // 消息id，64位整型
	}

	LinkMessage struct {
		ToUserName   string // 开发者微信号
		FromUserName string // 发送方帐号（一个OpenID）
		CreateTime   int64  // 消息创建时间 （整型）
		MsgType      string // link
		Title        string // 消息标题
		Description  string // 消息描述
		Url          string // 消息链接
		MsgId        int64  // 消息id，64位整型
	}
)

// 处理数据结构，xml字段
func Value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

// 新建一个文本回复信息，返回结构格式
func NewMsgTextResponse(ToUserName, FromUserName, Content string) *MsgTextResponse {
	textResp := &MsgTextResponse{
		CreateTime: time.Now().Unix(),
	}
	textResp.MsgType = Value2CDATA(MSG_TYPE_TEXT)
	textResp.ToUserName = Value2CDATA(ToUserName)
	textResp.FromUserName = Value2CDATA(FromUserName)
	textResp.Content = Value2CDATA(Content)
	return textResp
}

// 新建一个文本回复信息，返回二进制数组格式
func NewMsgTextResponseByte(ToUserName, FromUserName, Content string) ([]byte, error) {
	textResp := NewMsgTextResponse(ToUserName, FromUserName, Content)
	textRespByte, err := xml.Marshal(textResp)
	if err != nil {
		return nil, fmt.Errorf("NewMsgTextResponseByte xml.Marshal(textResp) error: %+v\n", err)
	}
	return textRespByte, nil
}

// 新建一个文本回复信息，返回字符串格式
func NewMsgTextResponseString(ToUserName, FromUserName, Content string) (string, error) {
	textRespByte, err := NewMsgTextResponseByte(ToUserName, FromUserName, Content)
	if err != nil {
		return "", fmt.Errorf("NewMsgTextResponseString error: %+v\n", err)
	}
	return string(textRespByte), nil
}

// 新建一个图片回复信息，返回结构格式
func NewMsgImageResponse(ToUserName, FromUserName, MediaId string) *MsgImageResponse {
	imageResp := &MsgImageResponse{
		CreateTime: time.Now().Unix(),
	}
	imageResp.MsgType = Value2CDATA(MSG_TYPE_IMG)
	imageResp.ToUserName = Value2CDATA(ToUserName)
	imageResp.FromUserName = Value2CDATA(FromUserName)
	imageResp.Image.MediaId = Value2CDATA(MediaId)
	return imageResp
}

// 新建一个图片回复信息，返回字节数组格式
func NewMsgImageResponseByte(ToUserName, FromUserName, MediaId string) ([]byte, error) {
	imageResp := NewMsgImageResponse(ToUserName, FromUserName, MediaId)
	imageRespByte, err := xml.Marshal(imageResp)
	if err != nil {
		return nil, fmt.Errorf("NewMsgImageResponseByte xml.Marshal(imageResp) error: %+v\n", err)
	}
	return imageRespByte, nil
}

// 新建一个图片回复信息，返回字符串格式
func NewMsgImageResponseString(ToUserName, FromUserName, MediaId string) (string, error) {
	imageRespByte, err := NewMsgImageResponseByte(ToUserName, FromUserName, MediaId)
	if err != nil {
		return "", fmt.Errorf("NewMsgImageResponseString error: %+v\n", err)
	}
	return string(imageRespByte), nil
}

// 新建一个语音回复信息，返回结构格式
func NewMsgVoiceResponse(ToUserName, FromUserName, MediaId string) *MsgVoiceResponse {
	voiceResp := &MsgVoiceResponse{
		CreateTime: time.Now().Unix(),
	}
	voiceResp.MsgType = Value2CDATA(MSG_TYPE_VOICE)
	voiceResp.ToUserName = Value2CDATA(ToUserName)
	voiceResp.FromUserName = Value2CDATA(FromUserName)
	voiceResp.Voice.MediaId = Value2CDATA(MediaId)
	return voiceResp
}

// 新建一个语音回复信息，返回字节数组格式
func NewMsgVoiceResponseByte(ToUserName, FromUserName, MediaId string) ([]byte, error) {
	voiceResp := NewMsgVoiceResponse(ToUserName, FromUserName, MediaId)
	voiceRespByte, err := xml.Marshal(voiceResp)
	if err != nil {
		return nil, fmt.Errorf("NewMsgVoiceResponseByte xml.Marshal(voiceResp) error: %+v\n", err)
	}
	return voiceRespByte, nil
}

// 新建一个语音回复信息，返回字符串格式
func NewMsgVoiceResponseString(ToUserName, FromUserName, MediaId string) (string, error) {
	voiceRespByte, err := NewMsgVoiceResponseByte(ToUserName, FromUserName, MediaId)
	if err != nil {
		return "", fmt.Errorf("NewMsgVoiceResponseString error: %+v\n", err)
	}
	return string(voiceRespByte), nil
}

// 新建一个视频回复信息，返回结构格式
func NewMsgVideoResponse(ToUserName, FromUserName, MediaId, Title, Description string) *MsgVideoResponse {
	videoResp := &MsgVideoResponse{
		CreateTime: time.Now().Unix(),
	}
	videoResp.MsgType = Value2CDATA(MSG_TYPE_VIDEO)
	videoResp.ToUserName = Value2CDATA(ToUserName)
	videoResp.FromUserName = Value2CDATA(FromUserName)
	videoResp.Video.MediaId = Value2CDATA(MediaId)
	videoResp.Video.Title = Value2CDATA(Title)
	videoResp.Video.Description = Value2CDATA(Description)
	return videoResp
}

// 新建一个视频回复信息，返回字节数组格式
func NewMsgVideoResponseByte(ToUserName, FromUserName, MediaId, Title, Description string) ([]byte, error) {
	videoResp := NewMsgVideoResponse(ToUserName, FromUserName, MediaId, Title, Description)
	videoRespByte, err := xml.Marshal(videoResp)
	if err != nil {
		return nil, fmt.Errorf("NewMsgVideoResponseByte xml.Marshal(videoResp) error: %+v\n", err)
	}
	return videoRespByte, nil
}

// 新建一个视频回复信息，返回字符串格式
func NewMsgVideoResponseString(ToUserName, FromUserName, MediaId, Title, Description string) (string, error) {
	videoRespByte, err := NewMsgVideoResponseByte(ToUserName, FromUserName, MediaId, Title, Description)
	if err != nil {
		return "", fmt.Errorf("NewMsgVideoResponseString error: %+v\n", err)
	}
	return string(videoRespByte), nil
}

// 新建一个音乐回复信息，返回结构格式
func NewMsgMusicResponse(ToUserName, FromUserName, Title, Description, ThumbMediaId, MusicUrl, HQMusicUrl string) *MsgMusicResponse {
	musicResp := &MsgMusicResponse{
		CreateTime: time.Now().Unix(),
	}
	musicResp.MsgType = Value2CDATA(MSG_TYPE_MUSIC)
	musicResp.ToUserName = Value2CDATA(ToUserName)
	musicResp.FromUserName = Value2CDATA(FromUserName)
	musicResp.Music.Title = Value2CDATA(Title)
	musicResp.Music.Description = Value2CDATA(Description)
	musicResp.Music.ThumbMediaId = Value2CDATA(ThumbMediaId)
	musicResp.Music.MusicUrl = Value2CDATA(MusicUrl)
	musicResp.Music.HQMusicUrl = Value2CDATA(HQMusicUrl)
	return musicResp
}

// 新建一个音乐回复信息，返回字节数组格式
func NewMsgMusicResponseByte(ToUserName, FromUserName, Title, Description, ThumbMediaId, MusicUrl, HQMusicUrl string) ([]byte, error) {
	musicResp := NewMsgMusicResponse(ToUserName, FromUserName, Title, Description, ThumbMediaId, MusicUrl, HQMusicUrl)
	musicRespByte, err := xml.Marshal(musicResp)
	if err != nil {
		return nil, fmt.Errorf("NewMsgMusicResponseByte xml.Marshal(musicResp) error: %+v\n", err)
	}
	return musicRespByte, nil
}

// 新建一个音乐回复信息，返回字符串格式
func NewMsgMusicResponseString(ToUserName, FromUserName, Title, Description, ThumbMediaId, MusicUrl, HQMusicUrl string) (string, error) {
	musicRespByte, err := NewMsgMusicResponseByte(ToUserName, FromUserName, Title, Description, ThumbMediaId, MusicUrl, HQMusicUrl)
	if err != nil {
		return "", fmt.Errorf("NewMsgMusicResponseString error: %+v\n", err)
	}
	return string(musicRespByte), nil
}

// 新建一条图文消息，多个图文结构可以组成一个图文信息, 8条以内
func NewMsgNewsItemResponse(Title, Description, PicUrl, Url string) *MsgNewsItemResponse {
	newsItemResp := &MsgNewsItemResponse{}
	newsItemResp.Title = Value2CDATA(Title)
	newsItemResp.Description = Value2CDATA(Description)
	newsItemResp.PicUrl = Value2CDATA(PicUrl)
	newsItemResp.Url = Value2CDATA(Url)
	return newsItemResp
}

// 新建一个图文回复信息，返回结构格式，items 个数限制为8条以内
func NewMsgNewsResponse(ToUserName, FromUserName string, items []MsgNewsItemResponse) *MsgNewsResponse {
	newsResp := &MsgNewsResponse{
		CreateTime: time.Now().Unix(),
	}
	newsResp.MsgType = Value2CDATA(MSG_TYPE_NEWS)
	newsResp.ToUserName = Value2CDATA(ToUserName)
	newsResp.FromUserName = Value2CDATA(FromUserName)
	newsResp.ArticleCount = len(items)
	newsResp.Articles = items
	return newsResp
}

// 新建一个图文回复信息，返回字节数组格式
func NewMsgNewsResponseByte(ToUserName, FromUserName string, items []MsgNewsItemResponse) ([]byte, error) {
	newsResp := NewMsgNewsResponse(ToUserName, FromUserName, items)
	newsRespByte, err := xml.Marshal(newsResp)
	if err != nil {
		return nil, fmt.Errorf("NewMsgNewsResponseByte xml.Marshal(newsResp) error: %+v\n", err)
	}
	return newsRespByte, nil
}

// 新建一个图文回复信息，返回字符串格式
func NewMsgNewsResponseString(ToUserName, FromUserName string, items []MsgNewsItemResponse) (string, error) {
	newsRespByte, err := NewMsgNewsResponseByte(ToUserName, FromUserName, items)
	if err != nil {
		return "", fmt.Errorf("NewMsgNewsResponseString error: %+v\n", err)
	}
	return string(newsRespByte), nil
}

// 处理微信消息
func (wm *WechatMp) wechatMessageHandler(msg []byte) string {
	fmt.Fprintf(common.WechatLoggerWriter, "WechatRequest: %s\n", msg)

	var response string

	// 根据消息类型分别处理
	switch common.GetMsgTypeFromWechatMessage(msg) {
	case MSG_TYPE_EVENT:
		response = wm.wechatEventMessageHandler(msg)
	case MSG_TYPE_TEXT:
		if wm.TextMessageHandler != nil {
			var textMessage TextMessage
			if err := xml.Unmarshal(msg, &textMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, textMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.TextMessageHandler(&textMessage)
		}
	case MSG_TYPE_IMG:
		if wm.ImageMessageHandler != nil {
			var imageMessage ImageMessage
			if err := xml.Unmarshal(msg, &imageMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, imageMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.ImageMessageHandler(&imageMessage)
		}
	case MSG_TYPE_VOICE:
		if wm.VoiceMessageHandler != nil {
			var voiceMessage VoiceMessage
			if err := xml.Unmarshal(msg, &voiceMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, voiceMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.VoiceMessageHandler(&voiceMessage)
		}
	case MSG_TYPE_VIDEO:
		if wm.VideoMessageHandler != nil {
			var videoMessage VideoMessage
			if err := xml.Unmarshal(msg, &videoMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, videoMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.VideoMessageHandler(&videoMessage)
		}
	case MSG_TYPE_VIDEO_SHORT:
		if wm.ShortVideoMessageHandler != nil {
			var shortVideoMessage ShortVideoMessage
			if err := xml.Unmarshal(msg, &shortVideoMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, shortVideoMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.ShortVideoMessageHandler(&shortVideoMessage)
		}
	case MSG_TYPE_LOCATION:
		if wm.LocationMessageHandler != nil {
			var locationMessage LocationMessage
			if err := xml.Unmarshal(msg, &locationMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, locationMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.LocationMessageHandler(&locationMessage)
		}
	case MSG_TYPE_LINK:
		if wm.LinkMessageHandler != nil {
			var linkMessage LinkMessage
			if err := xml.Unmarshal(msg, &linkMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, linkMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.LinkMessageHandler(&linkMessage)
		}
	default:
		response = RESPONSE_STRING_SUCCESS
	}
	return response
}

// 处理微信事件消息
func (wm *WechatMp) wechatEventMessageHandler(msg []byte) string {

	var response string

	switch common.GetMsgEventFromWechatMessage(msg) {
	case MSG_EVENT_SUBSCRIBE:
		// 关注公众号
		if wm.SubscribeHandler != nil {
			var subscribeMessage MsgEventSubscribe
			if err := xml.Unmarshal(msg, &subscribeMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, subscribeMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.SubscribeHandler(&subscribeMessage)
		}
	case MSG_EVENT_UNSUBSCRIBE:
		// 取消关注公众号
		if wm.UnSubscribeHandler != nil {
			var unSubscribeMessage MsgEventSubscribe
			if err := xml.Unmarshal(msg, &unSubscribeMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, unSubscribeMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.UnSubscribeHandler(&unSubscribeMessage)
		}
	case MSG_EVENT_SCAN:
		// 扫码公众号二维码
		if wm.ScanHandler != nil {
			var scanMessage MsgEventScan
			if err := xml.Unmarshal(msg, &scanMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, scanMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.ScanHandler(&scanMessage)
		}
	case MSG_EVENT_LOCATION:
		// 上报地理位置
		if wm.LocationHandler != nil {
			var locationMessage MsgEventLocation
			if err := xml.Unmarshal(msg, &locationMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, locationMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.LocationHandler(&locationMessage)
		}
	case MSG_EVENT_CLICK:
		// 点击自定义菜单，拉取消息
		if wm.MenuClickHandler != nil {
			var menuClickMessage MsgEventMenuClick
			if err := xml.Unmarshal(msg, &menuClickMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, menuClickMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.MenuClickHandler(&menuClickMessage)
		}
	case MSG_EVENT_VIEW:
		// 点击自定义菜单，跳转链接
		if wm.MenuViewHandler != nil {
			var menuViewMessage MsgEventMenuView
			if err := xml.Unmarshal(msg, &menuViewMessage); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, menuClickMessage) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.MenuViewHandler(&menuViewMessage)
		}
	case MSG_EVENT_QUALIFICATION_VERIFY_SUCCESS:
		// 资质认证成功
		if wm.QualificationVerifySuccessHandler != nil {
			var qualificationVerifySuccess MsgEventQualificationVerifySuccess
			if err := xml.Unmarshal(msg, &qualificationVerifySuccess); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, qualificationVerifySuccess) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.QualificationVerifySuccessHandler(&qualificationVerifySuccess)
		}
	case MSG_EVENT_QUALIFICATION_VERIFY_FAIL:
		// 资质认证失败
		if wm.QualificationVerifyFailHandler != nil {
			var qualificationVerifyFail MsgEventQualificationVerifyFail
			if err := xml.Unmarshal(msg, &qualificationVerifyFail); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, qualificationVerifyFail) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.QualificationVerifyFailHandler(&qualificationVerifyFail)
		}
	case MSG_EVENT_NAMING_VERIFY_SUCCESS:
		// 名称认证成功
		if wm.NamingVerifySuccessHandler != nil {
			var namingVerifySuccess MsgEventNamingVerifySuccess
			if err := xml.Unmarshal(msg, &namingVerifySuccess); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, namingVerifySuccess) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.NamingVerifySuccessHandler(&namingVerifySuccess)
		}
	case MSG_EVENT_NAMING_VERIFY_FAIL:
		// 名称认证失败
		if wm.NamingVerifyFailHandler != nil {
			var namingVerifyFail MsgEventNamingVerifyFail
			if err := xml.Unmarshal(msg, &namingVerifyFail); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, namingVerifyFail) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.NamingVerifyFailHandler(&namingVerifyFail)
		}
	case MSG_EVENT_ANNUAL_RENEW:
		// 年审通知
		if wm.AnnualRenewHandler != nil {
			var annualRenew MsgEventAnnualRenew
			if err := xml.Unmarshal(msg, &annualRenew); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, annualRenew) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.AnnualRenewHandler(&annualRenew)
		}
	case MSG_EVENT_VERIFY_EXPIRED:
		// 认证过期失效通知
		if wm.VerifyExpiredHandler != nil {
			var verifyExpired MsgEventVerifyExpired
			if err := xml.Unmarshal(msg, &verifyExpired); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, verifyExpired) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.VerifyExpiredHandler(&verifyExpired)
		}
	case MSG_EVENT_TEMPLATESENDJOBFINISH:
		// 发送模板消息结果通知
		if wm.SendTemplateFinishHandler != nil {
			var templateSendFinish MsgEventTemplateSendFinish
			if err := xml.Unmarshal(msg, &templateSendFinish); err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "DecodeMsg xml.Unmarshal(message: %s, templateSendFinish) error: %+v\n", msg, err)
				return RESPONSE_STRING_FAIL
			}
			response = wm.SendTemplateFinishHandler(&templateSendFinish)
		}
	default:
		fmt.Fprintf(common.WechatLoggerWriter, "wechat message find not msgType: %s\n", msg)
		response = ""
	}
	return response
}

//////////////////////////////////////////////////////  模版消息接口 /////////////////////////////////////////////////////

/*
行业代码查询

主行业	副行业	代码
IT科技	互联网/电子商务	1
IT科技	IT软件与服务	2
IT科技	IT硬件与设备	3
IT科技	电子技术	4
IT科技	通信与运营商	5
IT科技	网络游戏	6
金融业	银行	7
金融业	基金理财信托	8
金融业	保险	9
餐饮	餐饮	10
酒店旅游	酒店	11
酒店旅游	旅游	12
运输与仓储	快递	13
运输与仓储	物流	14
运输与仓储	仓储	15
教育	培训	16
教育	院校	17
政府与公共事业	学术科研	18
政府与公共事业	交警	19
政府与公共事业	博物馆	20
政府与公共事业	公共事业非盈利机构	21
医药护理	医药医疗	22
医药护理	护理美容	23
医药护理	保健与卫生	24
交通工具	汽车相关	25
交通工具	摩托车相关	26
交通工具	火车相关	27
交通工具	飞机相关	28
房地产	建筑	29
房地产	物业	30
消费品	消费品	31
商业服务	法律	32
商业服务	会展	33
商业服务	中介服务	34
商业服务	认证	35
商业服务	审计	36
文体娱乐	传媒	37
文体娱乐	体育	38
文体娱乐	娱乐休闲	39
印刷	印刷	40
其它	其它	41
*/

const (
	WECHAT_TEMPLATE_SET_INDUSTRY_API = `https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s`         // 设置所属行业
	WECHAT_TEMPLATE_GET_INDUSTRY_API = `https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s`             // 获取设置的行业信息
	WECHAT_TEMPLATE_ADD_API          = `https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s`         // 从模板库中添加模板信息
	WECHAT_TEMPLATE_GET_ALL_API      = `https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s` // 获取已添加至帐号下所有模板列表
	WECHAT_TEMPLATE_DELETE_API       = `https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s`     // 删除模板
	WECHAT_TEMPLATE_SEND_API         = `https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s`             // 发送模板消息
)

type (
	TemplateIndustryInfo struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	}

	TemplateIndustry struct {
		PrimaryIndustry   TemplateIndustryInfo `json:"primary_industry"`   // 帐号设置的主营行业
		SecondaryIndustry TemplateIndustryInfo `json:"secondary_industry"` // 帐号设置的副营行业
	}

	Template struct {
		TemplateId      string `json:"template_id"`                // 模板ID
		Title           string `json:"title,omitempty"`            // 模板标题
		PrimaryIndustry string `json:"primary_industry,omitempty"` // 模板所属行业的一级行业
		DeputyIndustry  string `json:"deputy_industry,omitempty"`  // 模板所属行业的二级行业
		Content         string `json:"content,omitempty"`          // 模板内容
		Example         string `json:"example,omitempty"`          // 模板示例
		common.PublicError
	}

	Templates struct {
		TemplateList []Template `json:"template_list"`
	}

	SendTemplate struct {
		Touser      string `json:"touser"`        // 接收者openid
		TemplateId  string `json:"template_id"`   // 模板ID
		Url         string `json:"url,omitempty"` // 模板跳转链接
		Miniprogram struct {
			Appid    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
			Pagepath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
		} `json:"miniprogram"` // 跳小程序所需数据，不需跳小程序可不用传该数据
		Data *map[string]SendTemplateData `json:"data"` // 模板数据
	}
	// 注：url和miniprogram都是非必填字段，若都不传则模板无跳转；若都传，会优先跳转至小程序。开发者可根据实际需要选择其中一种跳转方式即可。当用户的微信客户端版本不支持跳小程序时，将会跳转至url。

	SendTemplateData struct {
		Value string `json:"value"` // 模板内容消息
		Color string `json:"color"` // 模板内容字体颜色，不填默认为黑色
	}
)

// 设置所属行业
func (wm *WechatMp) SetTemplateIndustries(accessToken string, industryIds [2]int) (*common.PublicError, error) {

	var params = make(map[string]int)
	params["industry_id1"] = industryIds[0]
	params["industry_id2"] = industryIds[1]

	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_TEMPLATE_SET_INDUSTRY_API, accessToken), &params)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "SetTemplateIndustries http post error: %+v\n", err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "SetTemplateIndustries json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &result, nil
}

// 获取设置的行业信息
func (wm *WechatMp) GetTemplateIndustries(accessToken string) (*TemplateIndustry, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_TEMPLATE_GET_INDUSTRY_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetTemplateIndustries http get error: %+v\n", err)
		return nil, err
	}
	var templateIndustry TemplateIndustry
	if err = json.Unmarshal(resp, &templateIndustry); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "getTemplateIndustries json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &templateIndustry, nil
}

// 从模板库中添加模板
// templateCode: 模板库中模板的编号，有“TM**”和“OPENTMTM**”等形式
//TODO 待验证
func (wm *WechatMp) AddTemplate(accessToken, templateCode string) (*Template, error) {
	var params = make(map[string]string)
	params["template_id_short"] = templateCode
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_TEMPLATE_ADD_API, accessToken), &params)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddTemplate http post error: %+v\n", err)
		return nil, err
	}
	var template Template
	if err = json.Unmarshal(resp, &template); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddTemplate json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &template, nil
}

// 获取已添加至帐号下所有模板列表
func (wm *WechatMp) GetAllTemplates(accessToken string) (*Templates, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_TEMPLATE_GET_ALL_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetAllTemplates http get error: %+v\n", err)
		return nil, err
	}
	var templates Templates
	if err = json.Unmarshal(resp, &templates); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetAllTemplates json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &templates, nil
}

// 删除模板
func (wm *WechatMp) DeleteTemplate(accessToken, templateId string) (*common.PublicError, error) {
	var params = make(map[string]string)
	params["template_id"] = templateId
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_TEMPLATE_DELETE_API, accessToken), &params)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "DeleteTemplate http post error: %+v\n", err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "DeleteTemplate json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &result, nil
}

// 发送模板消息
func (wm *WechatMp) SendTemplate(accessToken string, template *SendTemplate) (*int64, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_TEMPLATE_SEND_API, accessToken), &template)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "SendTemplate http post error: %+v\n", err)
		return nil, err
	}

	// 发送模板消息的返回值
	type SendTemplateResponse struct {
		common.PublicError
		Msgid int64 `json:"msgid"`
	}
	var sendTemplateResponse SendTemplateResponse
	if err = json.Unmarshal(resp, &sendTemplateResponse); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "SendTemplate json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	var msgid = sendTemplateResponse.Msgid
	return &msgid, err
}

///////////////////////////////////////////////////////////  一次性订阅消息 //////////////////////////////////////////
const (
	WECHAT_SUBSCRIBEMSGACTION_API = `https://mp.weixin.qq.com/mp/subscribemsg?action=%s&appid=%s&scene=%d&template_id=%s&redirect_url=%s&reserved=%s#wechat_redirect`
	WECHAT_SUBSCRIBE_TEMPLATE_API = `https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=%s`
)

type SubscribeTemplate struct {
	Touser      string `json:"touser"`      // 填接收消息的用户openid
	TemplateId  string `json:"template_id"` // 订阅消息模板ID
	Url         string `json:"url"`         // 点击消息跳转的链接，需要有ICP备案
	Miniprogram struct {
		Appid    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，并且小程序要求是已发布的）
		Pagepath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Scene int                              `json:"scene,string"` // 订阅场景值
	Title string                           `json:"title"`        // 消息标题，15字以内
	Data  map[string]SubscribeTemplateData `json:"data"`         // 消息正文，value为消息内容文本（200字以内），没有固定格式，可用\n换行，color为整段消息内容的字体颜色（目前仅支持整段消息为一种颜色）
}

type SubscribeTemplateData struct {
	Value string `json:"value"` // 模板内容消息
	Color string `json:"color"` // 模板内容字体颜色，不填默认为黑色
}

// 注：url和miniprogram都是非必填字段，若都不传则模板无跳转；若都传，会优先跳转至小程序。开发者可根据实际需要选择其中一种跳转方式即可。当用户的微信客户端版本不支持跳小程序时，将会跳转至url。

// 一次性订阅消息的授权URL
func (wm *WechatMp) BuildSubscribeMsgURL(scene int, templateId, redirectUrl, reserved string) string {
	return fmt.Sprintf(WECHAT_SUBSCRIBEMSGACTION_API, "get_confirm", wm.Configure.AppId, scene, templateId, redirectUrl, reserved)
}

// 通过API推送订阅模板消息给到授权微信用户
//TODO 待验证
func (wm *WechatMp) SubscribeTemplate(accessToken string, template *SubscribeTemplate) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_SUBSCRIBE_TEMPLATE_API, accessToken), &template)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "SubscribeTemplate http post error: %+v\n", err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "SubscribeTemplate json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &result, nil
}
