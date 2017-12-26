package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
)

const (
	WECHAT_QRCODE_CREATE_API = `https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s` // 创建二维码API
	WECHAT_QRCODE_SHOW_API   = `https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s`           // 显示二维码API
)

const (
	WECHAT_QRCODE_SCENE           = "QR_SCENE"           // 二维码类型-临时的整型参数值
	WECHAT_QRCODE_STR_SCENE       = "QR_STR_SCENE"       // 二维码类型-临时的字符串参数值
	WECHAT_QRCODE_LIMIT_SCENE     = "QR_LIMIT_SCENE"     // 二维码类型-永久的整型参数值
	WECHAT_QRCODE_LIMIT_STR_SCENE = "QR_LIMIT_STR_SCENE" // 二维码类型-永久的字符串参数值
)

type (
	QrcodeRequest struct {
		ExpireSeconds int64  `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为30秒。
		ActionName    string `json:"action_name"`    // 二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
		ActionInfo    struct {
			Scene struct {
				SceneId  int64  `json:"scene_id,omitempty"`  // 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
				SceneStr string `json:"scene_str,omitempty"` // 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
			} `json:"scene"`
		} `json:"action_info"` // 二维码详细信息
	}

	QrcodeResponse struct {
		Ticket        string `json:"ticket"`                   // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
		ExpireSeconds int64  `json:"expire_seconds,omitempty"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
		Url           string `json:"url"`                      // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
	}
)

// 创建临时二维码，整型场景值
func (wm *WechatMp) QrcodeIntCreate(accessToken string, sceneId, expireSeconds int64) (*QrcodeResponse, error) {
	qrcodeRequest := &QrcodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    WECHAT_QRCODE_SCENE,
	}
	qrcodeRequest.ActionInfo.Scene.SceneId = sceneId
	return qrcodeCreate(accessToken, qrcodeRequest)
}

// 创建临时二维码，字符串场景值
func (wm *WechatMp) QrcodeStrCreate(accessToken string, sceneStr string, expireSeconds int64) (*QrcodeResponse, error) {
	qrcodeRequest := &QrcodeRequest{
		ExpireSeconds: expireSeconds,
		ActionName:    WECHAT_QRCODE_STR_SCENE,
	}
	qrcodeRequest.ActionInfo.Scene.SceneStr = sceneStr
	return qrcodeCreate(accessToken, qrcodeRequest)
}

// 创建永久二维码，整型场景值
func (wm *WechatMp) QrcodeIntLimitCreate(accessToken string, sceneId int64) (*QrcodeResponse, error) {
	qrcodeRequest := &QrcodeRequest{
		ActionName: WECHAT_QRCODE_LIMIT_SCENE,
	}
	qrcodeRequest.ActionInfo.Scene.SceneId = sceneId
	return qrcodeCreate(accessToken, qrcodeRequest)
}

// 创建永久二维码，字符串场景值
func (wm *WechatMp) QrcodeStrLimitCreate(accessToken, sceneStr string) (*QrcodeResponse, error) {
	qrcodeRequest := &QrcodeRequest{
		ActionName: WECHAT_QRCODE_LIMIT_STR_SCENE,
	}
	qrcodeRequest.ActionInfo.Scene.SceneStr = sceneStr
	return qrcodeCreate(accessToken, qrcodeRequest)
}

// 创建二维码
func qrcodeCreate(accessToken string, qrcodeRequest *QrcodeRequest) (*QrcodeResponse, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_QRCODE_CREATE_API, accessToken), qrcodeRequest)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "QrcodeCreate http error: %+v\n", err)
		return nil, err
	}

	var qrcodeResp QrcodeResponse
	if err = json.Unmarshal(resp, &qrcodeResp); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "QrcodeCreate json.Unmarshal() error: %+v\n", err)
		return nil, err
	}
	return &qrcodeResp, nil
}

// 显示二维码图片
func (wm *WechatMp) QrcodeShow(ticket string) ([]byte, error) {
	return common.HTTPGet(fmt.Sprintf(WECHAT_QRCODE_SHOW_API, ticket))
}
