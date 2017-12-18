package mp

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

// 验证微信消息签名是否正确
// wechatSignature 微信签名信息, wechatParams 微信消息参数
func CheckWechatAuthSign(wechatSignature string, wechatParams ...string) bool {
	return wechatSignature == SignMsg(wechatParams...)
}

// 签名微信消息
// 返回加密后的字符串
func SignMsg(wechatParams ...string) string {
	return SortSha1Signature("", wechatParams...)
}

// 先排序，再用sha1加密
// sep: 排序后的分隔符， params: 需要加密的字符串数组
func SortSha1Signature(sep string, params ...string) string {
	// 排序
	sort.Strings(params)
	// 加密
	s := sha1.New()
	io.WriteString(s, strings.Join(params, sep))
	return fmt.Sprintf("%x", s.Sum(nil))
}
