package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
	"os"
)

const (
	WechatGetcallbackipApi = `https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s` // 获取微信服务器IP地址
)

type CallBackIP struct {
	IPList []string `json:"ip_list"` // 微信服务器IP地址列表
	*common.PublicError
}

// 获取微信服务器IP地址
func (wm *WechatMp) GetCallBackIP(accessToken string) (*CallBackIP, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatGetcallbackipApi, accessToken))
	if err != nil {
		fmt.Fprintf(os.Stderr, "wechat mp getcallbackip error: %+v\n", err)
		return nil, fmt.Errorf("wechat mp getcallbackip error: %+v", err)
	}
	var callBackIP CallBackIP
	err = json.Unmarshal([]byte(resp), &callBackIP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wechat mp getcallbackip error: %+v\n", err)
		return nil, fmt.Errorf("wechat mp getcallbackip error: %+v", err)
	}
	return &callBackIP, nil
}
