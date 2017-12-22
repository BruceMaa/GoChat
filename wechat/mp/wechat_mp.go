package mp

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"io/ioutil"
	"net/http"
)

const (
	WECHAT_REQUEST_ECHOSTR       = "echostr"       // 微信认证服务器请求参数:返回字符串
	WECHAT_REQUEST_TIMESTAMP     = "timestamp"     // 微信服务器请求参数：时间戳
	WECHAT_REQUEST_NONCE         = "nonce"         // 微信服务器请求参数：随机字符串
	WECHAT_REQUEST_SIGNATURE     = "signature"     // 微信服务器请求参数：签名字符串
	WECHAT_REQUEST_ENCRYPT_TYPE  = "encrypt_type"  // 微信服务器请求参数：加密方式
	WECHAT_REQUEST_MSG_SIGNATURE = "msg_signature" // 微信服务器请求参数：消息签名字符串

	RESPONSE_STRING_SUCCESS = "success"
	RESPONSE_STRING_FAIL    = "fail"
	RESPONSE_STRING_INVALID = "invalid wechat server"

	WECHAT_LANGUAGE_ZH_CN = "zh_CH" // 微信语言，简体中文
	WECHAT_LANGUAGE_ZH_TW = "zh_TW" // 微信语言，繁体中文
	WECHAT_LANGUAGE_EN    = "en"    // 微信语言，英文

	WECHAT_ENCRYPT_TYPE = "aes" // 微信消息加密方式
)

type (
	WechatMpConfig struct {
		AppId          string `json:"app_id"`                     // 公众号appId
		AppSecret      string `json:"app_secret"`                 // 公众号appSecret
		Token          string `json:"token"`                      // 公众号Token
		EncodingAESKey string `json:"encoding_aes_key,omitempty"` // 公众号EncodingAESKey
	}

	WechatMp struct {
		Configure                         WechatMpConfig
		AccessToken                       *WechatAccessToken                    // 保存微信accessToken
		AccessTokenHandler                AccessTokenHandlerFunc                // 处理微信accessToken，如果有缓存，可以将accessToken存储到缓存中，默认存储到内存中
		SubscribeHandler                  SubscribeHandlerFunc                  // 关注微信公众号处理方法
		UnSubscribeHandler                UnSubscribeHandlerFunc                // 取消关注公众号处理方法
		ScanHandler                       ScanHandlerFunc                       // 扫描此微信公众号生成的二维码处理方法
		LocationHandler                   LocationHandlerFunc                   // 上报地理位置的处理方法
		MenuClickHandler                  MenuClickHandlerFunc                  // 自定义菜单点击的处理方法
		MenuViewHandler                   MenuViewHandlerFunc                   // 自定义菜单跳转外链的处理方法
		QualificationVerifySuccessHandler QualificationVerifySuccessHandlerFunc // 资质认证成功处理方法
		QualificationVerifyFailHandler    QualificationVerifyFailHandlerFunc    // 资质认证失败处理方法
		NamingVerifySuccessHandler        NamingVerifySuccessHandlerFunc        // 名称认证成功的处理方法
		NamingVerifyFailHandler           NamingVerifyFailHandlerFunc           // 名称认证失败的处理方法
		AnnualRenewHandler                AnnualRenewHandlerFunc                // 年审通知的处理方法
		VerifyExpiredHandler              VerifyExpireHandlerFunc               // 认证过期失效通知的处理方法
		SendTemplateFinishHandler         SendTemplateFinishHandlerFunc         // 发送模板消息结果通知
		TextMessageHandler                TextMessageHandlerFunc                // 发送文本信息的处理方法
		ImageMessageHandler               ImageMessageHandlerFunc               // 发送图片消息的处理方法
		VoiceMessageHandler               VoiceMessageHandlerFunc               // 发送语言消息的处理方法
		VideoMessageHandler               VideoMessageHandlerFunc               // 发送视频消息的处理方法
		ShortVideoMessageHandler          ShortVideoMessageHandlerFunc          // 发送短视频消息的处理方法
		LocationMessageHandler            LocationMessageHandlerFunc            // 上报地理位置的处理方法
		LinkMessageHandler                LinkMessageHandlerFunc                // 发送链接消息的处理方法
	}
)

// 新建一个微信公众号
func New(wechatMpConfig *WechatMpConfig) *WechatMp {
	var wechatMp = &WechatMp{}
	wechatMp.Configure = *wechatMpConfig
	wechatMp.SetAccessTokenHandlerFunc(WechatMpDefaultAccessTokenHandlerFunc)
	return wechatMp
}

// 用户在设置微信公众号服务器配置，并开启后，微信会发送一次认证请求，此函数即做此验证用
func (wm *WechatMp) AuthWechatServer(r *http.Request) string {
	echostr := r.FormValue(WECHAT_REQUEST_ECHOSTR)
	if wm.checkWechatSource(r) {
		return echostr
	}
	return RESPONSE_STRING_INVALID
}

// 检验认证来源是否为微信
func (wm *WechatMp) checkWechatSource(r *http.Request) bool {
	timestamp := r.FormValue(WECHAT_REQUEST_TIMESTAMP)
	nonce := r.FormValue(WECHAT_REQUEST_NONCE)
	signature := r.FormValue(WECHAT_REQUEST_SIGNATURE)
	return CheckWechatAuthSign(signature, wm.Configure.Token, timestamp, nonce)
}

// 检验消息来源，并且提取消息
func (wm *WechatMp) checkMessageSource(r *http.Request) (bool, []byte) {
	//openid := r.FormValue("openid") // openid，暂时还没想到为什么传值过来
	timestamp := r.FormValue(WECHAT_REQUEST_TIMESTAMP)
	nonce := r.FormValue(WECHAT_REQUEST_NONCE)

	// 读取request body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "checkMessageSource ioutil.ReadAll(r.Body) error: %+v\n", err)
		return false, nil
	}
	// 判断消息是否经过加密
	encrypt_type := r.FormValue(WECHAT_REQUEST_ENCRYPT_TYPE)
	if encrypt_type == WECHAT_ENCRYPT_TYPE {
		// 如果消息已经加密
		msg_signature := r.FormValue(WECHAT_REQUEST_MSG_SIGNATURE)
		var msgEncryptBody MsgEncryptBody
		if err = xml.Unmarshal(body, &msgEncryptBody); err != nil {
			fmt.Fprintf(common.WechatErrorLoggerWriter, "checkMessageSource xml.Unmarshal(body, &msgEncryptBody) error: %+v\n", err)
			return false, nil
		}
		check := CheckWechatAuthSign(msg_signature, timestamp, nonce, wm.Configure.Token, msgEncryptBody.Encrypt)
		var message []byte
		if check {
			// 验证成功，解密消息，返回正文的二进制数组格式
			message, err = wm.aesDecryptMessage(msgEncryptBody.Encrypt)
			if err != nil {
				fmt.Fprintf(common.WechatErrorLoggerWriter, "checkMessageSource wm.aesDecryptMessage(msgEncryptBody.Encrypt) error: %+v\n", err)
				return false, nil
			}
		}

		return check, message
	}
	// 如果消息未加密
	signature := r.FormValue(WECHAT_REQUEST_SIGNATURE)
	return CheckWechatAuthSign(signature, wm.Configure.Token, timestamp, nonce), body
}

// 加密后的微信消息结构
type MsgEncryptBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

// 加密发送消息
func (wm *WechatMp) aesEncryptMessage() {

}

// 解密收到的消息
func (wm *WechatMp) aesDecryptMessage(cipherMessage string) ([]byte, error) {
	// 先将密钥以及加密数据，做base64解码
	// = 为占位符
	aesKey, err := base64.StdEncoding.DecodeString(wm.Configure.EncodingAESKey + "=")
	if err != nil {
		return nil, fmt.Errorf("aesDecryptMessage base64 decode EncodingAESKey error: %+v\n", err)
	}
	message, err := base64.StdEncoding.DecodeString(cipherMessage)
	if err != nil {
		return nil, fmt.Errorf("aesDecryptMessage base64 decode encryptMessage error: %+v\n", err)
	}
	message, err = AESDecrypt(message, aesKey)
	if err != nil {
		return nil, fmt.Errorf("aesDecryptMessage AESDecrypt error: %+v\n", err)
	}

	// 解密完成后，提取正文
	return wm.extractDecryptMessage(message)
}

// 从解密后的消息中，提取正文msg
// msg_encrypt = Base64_Encode(AES_Encrypt[random(16B) + msg_len(4B) + msg + $AppID])
func (wm *WechatMp) extractDecryptMessage(plainData []byte) ([]byte, error) {
	// 前16位是随机字符, 直接跳过，17至20是正文的长度，先读取正文的长度
	buf := bytes.NewBuffer(plainData[16:20])
	var msgLength int32
	err := binary.Read(buf, binary.BigEndian, &msgLength)
	if err != nil {
		return nil, fmt.Errorf("extractDecryptMessage binary.Read(msgLength) error: %+v\n", err)
	}

	// 正文之后是appid， 可以再次验证，计算appid的起始位置
	appIdStart := msgLength + 20
	// 获取appid,并进行验证
	appId := string(plainData[appIdStart:])
	if wm.Configure.AppId != appId {
		// 验证消息中的appid未通过
		return nil, fmt.Errorf("local appid (%s) is not equal of message appid (%s)\n", wm.Configure.AppId, appId)
	}

	return plainData[20:appIdStart], nil
}

// 微信服务推送消息接收方法
func (wm *WechatMp) CallBackFunc(r *http.Request) string {
	// 首先，验证消息是否从微信服务发出
	valid, body := wm.checkMessageSource(r)
	if !valid {
		fmt.Fprintln(common.WechatErrorLoggerWriter, RESPONSE_STRING_INVALID)
		return RESPONSE_STRING_FAIL
	}
	return wm.wechatMessageHandler(body)
}

// 设置全局获取微信accessToken的方法
func (wm *WechatMp) SetAccessTokenHandlerFunc(handlerFunc AccessTokenHandlerFunc) {
	wm.AccessTokenHandler = handlerFunc
}

// 设置处理关注事件的方法
func (wm *WechatMp) SetSubscribeHandlerFunc(handlerFunc SubscribeHandlerFunc) {
	wm.SubscribeHandler = handlerFunc
}

// 设置处理取消关注事件的方法
func (wm *WechatMp) SetUnSubscribeHandlerFunc(handlerFunc UnSubscribeHandlerFunc) {
	wm.UnSubscribeHandler = handlerFunc
}

// 设置处理扫描事件的方法
func (wm *WechatMp) SetScanHandlerFunc(handlerFunc ScanHandlerFunc) {
	wm.ScanHandler = handlerFunc
}

// 设置处理上报地理位置的方法
func (wm *WechatMp) SetLocationHandlerFunc(handlerFunc LocationHandlerFunc) {
	wm.LocationHandler = handlerFunc
}

// 设置处理自定义菜单点击事件的方法
func (wm *WechatMp) SetMenuClickHandlerFunc(handlerFunc MenuClickHandlerFunc) {
	wm.MenuClickHandler = handlerFunc
}

// 设置处理自定义菜单跳转外链事件的方法
func (wm *WechatMp) SetMenuViewHandlerFunc(handlerFunc MenuViewHandlerFunc) {
	wm.MenuViewHandler = handlerFunc
}

// 设置处理微信text消息事件方法
func (wm *WechatMp) SetTextHandlerFunc(handlerFunc TextMessageHandlerFunc) {
	wm.TextMessageHandler = handlerFunc
}

// 设置处理微信image消息事件方法
func (wm *WechatMp) SetImageHandlerFunc(handlerFunc ImageMessageHandlerFunc) {
	wm.ImageMessageHandler = handlerFunc
}

// 设置处理微信voice消息事件方法
func (wm *WechatMp) SetVoiceHandlerFunc(handlerFunc VoiceMessageHandlerFunc) {
	wm.VoiceMessageHandler = handlerFunc
}

// 设置处理微信video消息事件方法
func (wm *WechatMp) SetVideoHandlerFunc(handlerFunc VideoMessageHandlerFunc) {
	wm.VideoMessageHandler = handlerFunc
}
