package mp

import (
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"strconv"
)

// 获取jsapi的签名
// timeStamp: 时间戳，nonceStr: 随机字符串，url: 调用JSSDK的路径
func (wm *WechatMp) JsapiSign(accessToken string, timeStamp int64, nonceStr, url string) (*string, error) {
	timestamp_str := "timestamp=" + strconv.FormatInt(timeStamp, 10)
	nonce_str := "noncestr=" + nonceStr
	url_str := "url=" + url
	jsapi_ticket, err := wm.GetTicket(accessToken, WECHAT_TICKET_TYPE_JSAPI)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "JsapiSign GetTicket error: %+v\n", err)
		return nil, err
	}
	jsapi_ticket_str := "jsapi_ticket=" + jsapi_ticket.Ticket
	jsapi_sign := SortSha1Signature("&", timestamp_str, nonce_str, url_str, jsapi_ticket_str)
	return &jsapi_sign, nil
}
