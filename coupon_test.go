package top

import (
	"context"
	"net/http"
	"testing"
)

func TestTopClient_TbkCouponGet(t *testing.T) {
	type fields struct {
		appKey     string
		appSecret  string
		url        string
		signMethod SignMethod
		httpClient *http.Client
		logger     LoggerInterface
	}
	type args struct {
		ctx context.Context
		req *TbkCouponGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TbkCouponGetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test coupon get",
			fields: fields{
				appKey:    "",
				appSecret: "",
			},
			args: args{
				ctx: context.Background(),
				req: NewTbkCouponGetRequest("", "", 584399378309),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTopClient(tt.fields.appKey, tt.fields.appSecret)

			got, err := c.TbkCouponGet(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopClient.TbkCouponGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got %+v", got)
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopClient.TbkCouponGet() = %v, want %v", got, tt.want)
			}*/
		})
	}
}
