package top

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type TbkCouponGetRequest struct {
	Me         string `json:"me,omitempty"`
	ItemID     int64  `json:"item_id,omitempty"`
	ActivityID string `json:"activity_id,omitempty"`

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

func NewTbkCouponGetRequest(me, activityID string, itemID int64) *TbkCouponGetRequest {
	return &TbkCouponGetRequest{
		Me:         me,
		ItemID:     itemID,
		ActivityID: activityID,
		Method:     "taobao.tbk.coupon.get",

		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		APIVersion: APIVersion,
		Format:     "json",
	}
}

func (r *TbkCouponGetRequest) Signature(ctx context.Context, appkey, secret string, signMethod SignMethod) error {
	var (
		sign string
		err  error
	)
	r.SignMethod = signMethod.String()
	r.AppKey = appkey
	log.Printf("sign method:%s", signMethod.String())
	if signMethod == SignMethodMD5 {
		sign, err = signatureMD5(secret, r)
	} else if signMethod == SignMethodHMAC {
		sign, err = signatureHMAC(secret, r)
	} else {
		return errors.New("sign method not support")
	}
	if err != nil {
		return fmt.Errorf("signature error :%w", err)
	}
	r.Sign = sign
	return nil
}

type TbkCouponGetResponse struct {
	TbkCouponGetResponse TbkCouponGetResponseTbkCouponGetResponse `json:"tbk_coupon_get_response"`
	ErrorResponse        ErrorResponse                            `json:"error_response"`
}

type TbkCouponGetResponseData struct {
	CouponStartFee    string `json:"coupon_start_fee"`
	CouponRemainCount int    `json:"coupon_remain_count"`
	CouponTotalCount  int    `json:"coupon_total_count"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponAmount      string `json:"coupon_amount"`
	CouponSrcScene    int    `json:"coupon_src_scene"`
	CouponType        int    `json:"coupon_type"`
	CouponActivityID  string `json:"coupon_activity_id"`
}

type TbkCouponGetResponseTbkCouponGetResponse struct {
	Data TbkCouponGetResponseData `json:"data"`
}

func (c *TopClient) TbkCouponGet(ctx context.Context, req *TbkCouponGetRequest) (*TbkCouponGetResponse, error) {
	respData, err := c.execute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("execute error %w", err)
	}
	resp := &TbkCouponGetResponse{}
	err = json.Unmarshal(respData, resp)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error %w", err)
	}

	if resp.ErrorResponse.Code > 0 {

	}

	return resp, nil
}
