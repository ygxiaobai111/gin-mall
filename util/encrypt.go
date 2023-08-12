package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
)

var Encrypt *Encryption

// AES对称加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEnCryption()
}
func NewEnCryption() *Encryption {
	return &Encryption{}
}

// 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// AesEncoding加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		fmt.Println("加密失败")
		return ""
	}
	//密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize()) //因字节长度不够，所以需要进行字节填充
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// 去除填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	//去掉的长度
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:(len(dst) - unpadNum)]
	return str, nil
}

// AesDecoding 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		fmt.Println("解密失败")
		return ""
	}
	dst := make([]byte, len(pwdByte))
	dst, err = UnPadPwd(dst) //去除填充
	if err != nil {
		fmt.Println("填充去除失败")
		return "0"
	}
	return string(dst)

}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
