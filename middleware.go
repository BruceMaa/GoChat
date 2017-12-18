package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)


func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		switch r.Method {
		case "POST":
			r.ParseForm()
			requestBody, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				Logger.Errorf("Path:%+v | Method:%+v | Params:%+v | Body:%+v", r.URL, r.Method, r.Form, string(requestBody))
			} else {
				Logger.Debugf("Path:%+v | Method:%+v | Params:%+v | Body:%+v", r.URL, r.Method, r.Form, string(requestBody))
			}
		case "GET":
			fallthrough
		default:
			Logger.Debugf("Path:%+v | Method:%+v | Params:%+v", r.URL, r.Method, r.URL.Query())
		}
		c.Next()
	}
}
