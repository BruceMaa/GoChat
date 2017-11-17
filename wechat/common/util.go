package common

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

// HTTP Get 请求
func HTTPGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: url=%v, statusCode=%v", url, resp.StatusCode)
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
		return nil, fmt.Errorf("http POST Json error: url=%v, statusCode=%v", url, resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// HTTP Post 表单提交请求
func HTTPPostForm(url string, params map[string][]string) ([]byte, error) {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil , fmt.Errorf("http POST Form error: url=%v, statusCode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
