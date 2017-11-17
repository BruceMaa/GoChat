package mp

type (
	WechatMpConfig struct {
		AppId          string `json:"app_id"`                     // 公众号appId
		AppSecret      string `json:"app_secret"`                 // 公众号appSecret
		Token          string `json:"token"`                      // 公众号Token
		EncodingAESKey string `json:"encoding_aes_key,omitempty"` // 公众号EncodingAESKey
	}

	WechatMp struct {
		Configure WechatMpConfig
	}
)
