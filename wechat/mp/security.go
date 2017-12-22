package mp

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

// 验证微信消息签名是否正确
// wechatSignature 微信签名信息, wechatParams 微信消息参数
func CheckWechatAuthSign(wechatSignature string, wechatParams ...string) bool {
	return wechatSignature == SignMsg(wechatParams...)
}

// 签名微信消息
// 返回加密后的字符串
func SignMsg(wechatParams ...string) string {
	return SortSha1Signature("", wechatParams...)
}

// 先排序，再用sha1加密
// sep: 排序后的分隔符， params: 需要加密的字符串数组
func SortSha1Signature(sep string, params ...string) string {
	// 排序
	sort.Strings(params)
	// 加密
	s := sha1.New()
	io.WriteString(s, strings.Join(params, sep))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// AES 解密
// cipherData: 待解密信息， aesKey: 信息加密时用到的key
func AESDecrypt(cipherData, aesKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, aesKey[:blockSize])
	data := make([]byte, len(cipherData))
	blockMode.CryptBlocks(data, cipherData)
	data = PKCS7UnPadding(data)
	return data, nil
}

// AES 加密
// plainData: 待加密信息， aesKey: 加密需要用到的key
func AESEncrypt(plainData, aesKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainData = PKCS7Padding(plainData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aesKey[:blockSize])
	data := make([]byte, len(plainData))
	blockMode.CryptBlocks(data, plainData)
	return data, nil
}

func PKCS7UnPadding(cipherData []byte) []byte {
	length := len(cipherData)
	unpadding := int(cipherData[length-1])
	return cipherData[:(length - unpadding)]
}

func PKCS7Padding(plainData []byte, blockSize int) []byte {
	padding := blockSize - len(plainData)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainData, padtext...)
}
