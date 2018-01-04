package mp

import (
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
	"strconv"
)

// 获取jsapi的签名
// timeStamp: 时间戳，nonceStr: 随机字符串，url: 调用JSSDK的路径
func (wm *WechatMp) JsapiSign(accessToken string, timeStamp int64, nonceStr, url string) (*string, error) {
	timestampStr := "timestamp=" + strconv.FormatInt(timeStamp, 10)
	nonce := "noncestr=" + nonceStr
	urlStr := "url=" + url
	jsapiTicket, err := wm.GetTicket(accessToken, WechatTicketTypeJsapi)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "JsapiSign GetTicket error: %+v\n", err)
		return nil, err
	}
	jsapiTicketStr := "jsapi_ticket=" + jsapiTicket.Ticket
	jsapiSign := SortSha1Signature("&", timestampStr, nonce, urlStr, jsapiTicketStr)
	return &jsapiSign, nil
}
