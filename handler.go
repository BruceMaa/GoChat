package main

import (
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/mp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 初始化配置微信
var wechatMp = &mp.WechatMp{}

func init() {
	wechatMpConfig := &ChatConfig.WechatConfig.WechatMpConfig
	wechatMp = mp.New(wechatMpConfig)
}

// 设置微信公众号服务
func wechatMpServer(router *gin.Engine) {
	wechatApis := router.Group("/wechat")
	{
		wechatApis.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, wechatMp.AuthWechatServer(c.Request))
		})

		wechatApis.POST("", func(c *gin.Context) {
			c.String(http.StatusOK, wechatMp.CallBackFunc(c.Request))
		})

		wechatApis.GET("/subscribe", func(c *gin.Context) {
			openid := c.Query("openid")
			templateid := c.Query("template_id")
			action := c.Query("action")
			scene := c.Query("scene")
			reserved := c.Query("reserved")
			fmt.Printf("openid: %s, templateid: %s, action: %s, scene: %s, reserved: %s\n", openid, templateid, action, scene, reserved)
		})
	}

	wechatMp.SetSubscribeHandlerFunc(WechatMpMsgEventSubscribeHandler)
	wechatMp.SetUnSubscribeHandlerFunc(WechatMpMsgEventUnSubscribeHandler)
	wechatMp.SetScanHandlerFunc(WechatMpMsgEventScanHandler)
	wechatMp.SetLocationHandlerFunc(WechatMpMsgEventLocationHandler)
	wechatMp.SetMenuClickHandlerFunc(WechatMpMsgEventMenuClickHandler)
	wechatMp.SetMenuViewHandlerFunc(WechatMpMsgEventMenuViewHandler)
}

func WechatMpMsgEventSubscribeHandler(subscribeMessage *mp.MsgEventSubscribe) string {
	fmt.Printf("WechatMpMsgEventSubscribeHandler subscribeMessage: %+v\n", subscribeMessage)
	textRespStr, err := mp.NewMsgTextResponseString(subscribeMessage.FromUserName, subscribeMessage.ToUserName, "谢谢关注")
	if err != nil {
		fmt.Printf("msgTextResponse error: %+v\n", err)
		return ""
	}
	fmt.Printf("WechatMpMsgEventSubscribeHandler textRespStr: %s\n", textRespStr)
	return textRespStr
}

func WechatMpMsgEventUnSubscribeHandler(unSubscribeMessage *mp.MsgEventSubscribe) string {
	fmt.Printf("WechatMpMsgEventUnSubscribeHandler unSubscribeMessage: %+v\n", unSubscribeMessage)
	return ""
}

func WechatMpMsgEventScanHandler(scanMessage *mp.MsgEventScan) string {
	fmt.Printf("WechatMpMsgEventScanHandler scanMessage: %+v\n", scanMessage)
	textRespStr, err := mp.NewMsgTextResponseString(scanMessage.FromUserName, scanMessage.ToUserName, "扫码事件")
	if err != nil {
		fmt.Printf("msgTextResponse error: %+v\n", err)
		return ""
	}
	fmt.Printf("WechatMpMsgEventScanHandler textRespStr: %s\n", textRespStr)
	return textRespStr
}

func WechatMpMsgEventLocationHandler(locationMessage *mp.MsgEventLocation) string {
	fmt.Printf("WechatMpMsgEventLocationHandler locationMessage: %+v\n", locationMessage)
	content := "上报地理位置事件，\n纬度：" + locationMessage.Latitude + "\n经度：" + locationMessage.Longitude + "\n精度：" + locationMessage.Precision
	textRespStr, err := mp.NewMsgTextResponseString(locationMessage.FromUserName, locationMessage.ToUserName, content)
	if err != nil {
		fmt.Printf("msgTextResponse error: %+v\n", err)
		return ""
	}
	fmt.Printf("WechatMpMsgEventLocationHandler textRespStr: %s\n", textRespStr)
	return textRespStr
}

func WechatMpMsgEventMenuClickHandler(menuClickMessage *mp.MsgEventMenuClick) string {
	fmt.Printf("WechatMpMsgEventMenuClickHandler menuClickMessage: %+v\n", menuClickMessage)
	textRespStr, err := mp.NewMsgTextResponseString(menuClickMessage.FromUserName, menuClickMessage.ToUserName, "菜单点击事件: "+menuClickMessage.EventKey)
	if err != nil {
		fmt.Printf("msgTextResponse error: %+v\n", err)
		return ""
	}
	fmt.Printf("WechatMpMsgEventMenuClickHandler textRespStr: %s\n", textRespStr)
	return textRespStr
}

func WechatMpMsgEventMenuViewHandler(menuViewMessage *mp.MsgEventMenuView) string {
	fmt.Printf("WechatMpMsgEventMenuViewHandler menuViewMessage: %+v\n", menuViewMessage)
	textRespStr, err := mp.NewMsgTextResponseString(menuViewMessage.FromUserName, menuViewMessage.ToUserName, "菜单外链事件: "+menuViewMessage.EventKey)
	if err != nil {
		fmt.Printf("msgTextResponse error: %+v\n", err)
		return ""
	}
	fmt.Printf("WechatMpMsgEventMenuViewHandler textRespStr: %s\n", textRespStr)
	return textRespStr
}

// 设置对内网API路由
func apisRouter(router *gin.Engine) {
	apis := router.Group("/apis")
	{
		apis.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"messages": "Welcome GoChat Server API",
			})
		})

		apis.GET("/qrcode/create", func(c *gin.Context) {
			param := c.Query("scene")
			fmt.Printf("/apis/qrcode/create param: %s\n", param)
			scene, _ := strconv.ParseInt(param, 10, 64)
			fmt.Printf("/apis/qrcode/create scene: %d\n", scene)

			qrcodeResponse, _ := WechatMpQrcodeCreate(scene)
			data, _ := wechatMp.QrcodeShow(qrcodeResponse.Ticket)
			c.Header("Content-Type", "image/jpg")
			c.Writer.Write(data)
		})

		apis.GET("/shorturl", func(c *gin.Context) {
			longUrl := c.Query("longUrl")
			fmt.Printf("/apis/shorturl param: %s\n", longUrl)
			shortUrlResponse, _ := wechatMp.Long2Short(wechatMp.AccessTokenHandler(wechatMp), longUrl)
			c.JSON(http.StatusOK, shortUrlResponse)
		})
	}
}

func WechatMpQrcodeCreate(scene int64) (*mp.QrcodeResponse, error) {
	token := wechatMp.AccessTokenHandler(wechatMp)
	qrcodeResponse, err := wechatMp.QrcodeIntCreate(token, scene, 60)
	if err != nil {
		return nil, err
	}
	return qrcodeResponse, nil
}
