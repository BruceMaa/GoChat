package main

import (
	"encoding/xml"
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/mp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	fmt.Printf("WechatMpMsgEventSubscribeHandler: %+v\n", subscribeMessage)
	msgTextResponse := &mp.MsgTextResponse{
		ToUserName:   subscribeMessage.FromUserName,
		FromUserName: subscribeMessage.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      mp.MSG_TYPE_TEXT,
		Content:      "谢谢关注",
	}
	data, err := xml.Marshal(msgTextResponse)
	if err != nil {
		fmt.Printf("msgTextResponse xml.Marshal error: %+v\n", err)
		return ""
	}
	fmt.Printf("msgTextResponse: %s", data)
	return string(data)
}

func WechatMpMsgEventUnSubscribeHandler(unSubscribeMessage *mp.MsgEventSubscribe) string {
	fmt.Printf("WechatMpMsgEventUnSubscribeHandler: %+v\n", unSubscribeMessage)
	return ""
}

func WechatMpMsgEventScanHandler(scanMessage *mp.MsgEventScan) string {
	fmt.Printf("WechatMpMsgEventScanHandler: %+v\n", scanMessage)
	msgTextResponse := &mp.MsgTextResponse{
		ToUserName:   scanMessage.FromUserName,
		FromUserName: scanMessage.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      mp.MSG_TYPE_TEXT,
		Content:      "扫码事件",
	}
	data, err := xml.Marshal(msgTextResponse)
	if err != nil {
		fmt.Printf("msgTextResponse xml.Marshal error: %+v\n", err)
		return ""
	}
	fmt.Printf("msgTextResponse: %s", data)
	return string(data)
}

func WechatMpMsgEventLocationHandler(locationMessage *mp.MsgEventLocation) string {
	fmt.Printf("WechatMpMsgEventLocationHandler: %+v\n", locationMessage)
	msgTextResponse := &mp.MsgTextResponse{
		ToUserName:   locationMessage.FromUserName,
		FromUserName: locationMessage.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      mp.MSG_TYPE_TEXT,
		Content:      "上报地理位置事件，\n纬度：" + locationMessage.Latitude + "\n经度：" + locationMessage.Longitude + "\n精度：" + locationMessage.Precision,
	}
	data, err := xml.Marshal(msgTextResponse)
	if err != nil {
		fmt.Printf("msgTextResponse xml.Marshal error: %+v\n", err)
		return ""
	}
	fmt.Printf("msgTextResponse: %s", data)
	return string(data)
}

func WechatMpMsgEventMenuClickHandler(menuClickMessage *mp.MsgEventMenuClick) string {
	fmt.Printf("WechatMpMsgEventMenuClickHandler: %+v\n", menuClickMessage)
	msgTextResponse := &mp.MsgTextResponse{
		ToUserName:   menuClickMessage.FromUserName,
		FromUserName: menuClickMessage.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      mp.MSG_TYPE_TEXT,
		Content:      "菜单点击事件: " + menuClickMessage.EventKey,
	}
	data, err := xml.Marshal(msgTextResponse)
	if err != nil {
		fmt.Printf("msgTextResponse xml.Marshal error: %+v\n", err)
		return ""
	}
	fmt.Printf("msgTextResponse: %s", data)
	return string(data)
}

func WechatMpMsgEventMenuViewHandler(menuViewMessage *mp.MsgEventMenuView) string {
	fmt.Printf("WechatMpMsgEventMenuViewHandler: %+v\n", menuViewMessage)
	msgTextResponse := &mp.MsgTextResponse{
		ToUserName:   menuViewMessage.FromUserName,
		FromUserName: menuViewMessage.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      mp.MSG_TYPE_TEXT,
		Content:      "菜单外链事件: " + menuViewMessage.EventKey,
	}
	data, err := xml.Marshal(msgTextResponse)
	if err != nil {
		fmt.Printf("msgTextResponse xml.Marshal error: %+v\n", err)
		return ""
	}
	fmt.Printf("msgTextResponse: %s", data)
	return string(data)
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
