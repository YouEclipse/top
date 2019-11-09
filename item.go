package top

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type TbkItemInfoGetRequest struct {
	NumIids  string `json:"num_iids"`
	Platform int64  `json:"platform"`
	IP       string `json:"ip,omitempty"`

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

func NewTbkItemInfoGetRequest(numIids, ip string, platform int64) *TbkItemInfoGetRequest {
	return &TbkItemInfoGetRequest{
		NumIids:  numIids,
		Platform: platform,
		IP:       ip,
		Method:   "taobao.tbk.item.info.get",

		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		APIVersion: APIVersion,
		Format:     "json",
	}
}

func (r *TbkItemInfoGetRequest) Signature(ctx context.Context, appkey, secret string, signMethod SignMethod) error {
	var (
		sign string
		err  error
	)
	r.SignMethod = signMethod.String()
	r.AppKey = appkey
	log.Printf("%s", signMethod)
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

func (r *TbkItemInfoGetRequest) GetRequestData(ctx context.Context) ([]byte, error) {
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

	return bytes.TrimRight(data.Bytes(), "&"), nil
}

type TbkItemInfoGetResponse struct {
	TbkItemInfoGetResponseItem TbkItemInfoGetResponseItem `json:"tbk_item_info_get_response"`
	ErrorResponse              ErrorResponse              `json:"error_response"`
}

type Results struct {
	NTbkItem []NTbkItem `json:"n_tbk_item"`
}

type TbkItemInfoGetResponseItem struct {
	Results Results `json:"results"`
}

type SmallImages struct {
	String []string `json:"string"`
}

type NTbkItem struct {
	CatName                    string      `json:"cat_name"`
	NumIid                     int         `json:"num_iid"`
	Title                      string      `json:"title"`
	PictURL                    string      `json:"pict_url"`
	SmallImages                SmallImages `json:"small_images"`
	ReservePrice               string      `json:"reserve_price"`
	ZkFinalPrice               string      `json:"zk_final_price"`
	UserType                   int         `json:"user_type"`
	Provcity                   string      `json:"provcity"`
	ItemURL                    string      `json:"item_url"`
	SellerID                   int         `json:"seller_id"`
	Volume                     int         `json:"volume"`
	Nick                       string      `json:"nick"`
	CatLeafName                string      `json:"cat_leaf_name"`
	IsPrepay                   bool        `json:"is_prepay"`
	ShopDsr                    int         `json:"shop_dsr"`
	Ratesum                    int         `json:"ratesum"`
	IRfdRate                   bool        `json:"i_rfd_rate"`
	HGoodRate                  bool        `json:"h_good_rate"`
	HPayRate30                 bool        `json:"h_pay_rate30"`
	FreeShipment               bool        `json:"free_shipment"`
	MaterialLibType            string      `json:"material_lib_type"`
	PresaleDiscountFeeText     string      `json:"presale_discount_fee_text"`
	PresaleTailEndTime         int64       `json:"presale_tail_end_time"`
	PresaleTailStartTime       int64       `json:"presale_tail_start_time"`
	PresaleEndTime             int64       `json:"presale_end_time"`
	PresaleStartTime           int64       `json:"presale_start_time"`
	PresaleDeposit             string      `json:"presale_deposit"`
	JuPlayEndTime              int64       `json:"ju_play_end_time"`
	JuPlayStartTime            int64       `json:"ju_play_start_time"`
	PlayInfo                   string      `json:"play_info"`
	TmallPlayActivityEndTime   int64       `json:"tmall_play_activity_end_time"`
	TmallPlayActivityStartTime int64       `json:"tmall_play_activity_start_time"`
}

func (c *TopClient) TbkItemInfoGet(ctx context.Context, req *TbkItemInfoGetRequest) (*TbkItemInfoGetResponse, error) {
	respData, err := c.execute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("execute error %w", err)
	}
	resp := &TbkItemInfoGetResponse{}
	err = json.Unmarshal(respData, resp)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error %w", err)
	}

	if resp.ErrorResponse.Code > 0 {

	}

	return resp, nil
}
