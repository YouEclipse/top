package top

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"log"
	"strings"
)

type SignMethod string

func (m SignMethod) String() string {
	return string(m)
}

const (
	SignMethodMD5  SignMethod = "md5"
	SignMethodHMAC SignMethod = "hmac"
)

type TopRequest interface {
	Signature(ctx context.Context, appkey, secret string, signMethod SignMethod) error
}
type ErrorResponse struct {
	SubMsg  string `json:"sub_msg"`
	Code    int    `json:"code"`
	SubCode string `json:"sub_code"`
	Msg     string `json:"msg"`
}

type CommonFields struct {
	// API接口名称 required
	Method string `json:"method,omitempty"`
	// TOP分配给应用的AppKey
	AppKey string `json:"app_key,omitempty"`
	// 用户登录授权成功后，TOP颁发给应用的授权信息 optional
	Session string `json:"session,omitempty"`
	// 时间戳 required
	Timestamp string `json:"timestamp,omitempty"`
	// 	响应格式 json,xml   optional 默认xml
	Format string `json:"format,omitempty"`
	// API协议版本 2.0
	APIVersion string `json:"v,omitempty"`
	// 合作伙伴身份标识 optional
	PartnerID string `json:"partner_id,omitempty"`
	// 被调用的目标AppKey optional
	TargetAppKey string `json:"target_app_key,omitempty"`
	// 是否采用精简JSON返回格式
	Simplify bool `json:"simplify,omitempty"`
	// sign_method 签名的摘要算法 hmac/md5
	SignMethod string `json:"sign_method,omitempty"`
	// API输入参数签名结果
	Sign string `json:"sign,omitempty"`
}

func signatureMD5(secret string, data TopRequest) (string, error) {
	pairs := newKVPairList()
	pairs.load(data)

	var (
		sign = bytes.NewBufferString(secret)
	)

	for _, pair := range pairs.list {
		if pair.key == "sign" {
			continue
		}
		sign.WriteString(pair.key)
		sign.WriteString(pair.value)
	}
	h := md5.New()
	h.Write(sign.Bytes())

	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
}

func signatureHMAC(secret string, data TopRequest) (string, error) {
	pairs := newKVPairList()
	pairs.load(data)

	var (
		sign = bytes.NewBuffer(nil)
	)

	for _, pair := range pairs.list {
		if pair.key == "sign" {
			continue
		}
		sign.WriteString(pair.key)
		sign.WriteString(pair.value)
	}
	log.Printf("%s", sign.String())
	h := hmac.New(md5.New, []byte(secret))
	h.Write(sign.Bytes())
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
}

func getRequestData(r interface{}) []byte {
	pairs := newKVPairList()
	pairs.load(r)
	var (
		data = bytes.NewBuffer(nil)
	)

	for _, pair := range pairs.list {
		data.WriteString(pair.key)
		data.WriteString("=")
		data.WriteString(pair.value)
		data.WriteString("&")
	}

	return bytes.TrimRight(data.Bytes(), "&")
}
