package mp

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

// 获取微信公众号的认证echo信息
// 用户在设置微信公众号服务器配置，并开启后，微信会发送一次认证请求，此函数即做此验证用
func (wm *WechatMp) AuthEchostr(token string, r *http.Request) string {
	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")
	echostr := r.FormValue("echostr")
	if CheckWecahtAuthSign(signature, token, timestamp, nonce) {
		return echostr
	}
	return ""
}

// 验证微信消息签名是否正确
// wechatSignature 微信签名信息, wechatParams 微信消息参数
func CheckWecahtAuthSign(wechatSignature string, wechatParams ...string) bool {
	return wechatSignature == SignMsg(wechatParams...)
}

// 签名微信消息
// 返回加密后的字符串
func SignMsg(wechatParams ...string) string {
	// 排序
	sort.Strings(wechatParams)
	// 加密
	s := sha1.New()
	io.WriteString(s, strings.Join(wechatParams, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
