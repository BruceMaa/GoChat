package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"regexp"
	"time"
)

// HTTP Get 请求
func HTTPGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: url=%v, statusCode=%v\n", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

// HTTP Post Json请求
func HTTPPostJson(url string, obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)

	body := bytes.NewBuffer(jsonData)
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http POST Json error: url=%v, statusCode=%v\n", url, resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// 文件上传结构体
type MultipartFormFile struct {
	FieldName string    // 文件入参名称
	FileName  string    // 文件名称
	Reader    io.Reader // 文件输入流
}

// HTTP POST form表单提交
// url:文件上传访问路径
// fields: 表单字段
// files: 表单文件
func HTTPPostForm(url string, fields map[string]string, files ...MultipartFormFile) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	for key, val := range fields {
		if err := writer.WriteField(key, val); err != nil {
			return nil, err
		}
	}
	for _, file := range files {
		formFile, err := writer.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return nil, err
		}

		// 从文件读取数据，写入表单
		if _, err = io.Copy(formFile, file.Reader); err != nil {
			return nil, err
		}
	}

	// 发送表单
	contentType := writer.FormDataContentType()
	// 发送之前必须调用Close()以写入结尾行
	if err := writer.Close(); err != nil {
		return nil, err
	}
	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

// 从微信消息中，提取MsgType
func GetMsgTypeFromWechatMessage(message []byte) string {
	regStr := `<MsgType><!\[CDATA\[(\w+)\]\]></MsgType>`
	msgReg := regexp.MustCompile(regStr)
	result := msgReg.FindSubmatch(message)
	if len(result) < 2 {
		// 没有找到匹配
		return ""
	}
	return string(result[1])
}

// 从微信事件消息中，提取事件类型
func GetMsgEventFromWechatMessage(message []byte) string {
	regStr := `<Event><!\[CDATA\[(\w+)\]\]></Event>`
	msgReg := regexp.MustCompile(regStr)
	result := msgReg.FindSubmatch(message)
	if len(result) < 2 {
		// 没有找到匹配
		return ""
	}
	return string(result[1])
}

//获取随机字节数组
func GetRandomString(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nonce := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, nonce[r.Intn(len(nonce))])
	}
	return result
}
