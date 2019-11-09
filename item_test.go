package top

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestTopClient_TbkItemInfoGet(t *testing.T) {
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
		req *TbkItemInfoGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TbkItemInfoGetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test item get",
			fields: fields{
				appKey:    "28111323",
				appSecret: "a3ff5428d6a32795a96732bb552cc802",
			},
			args: args{
				ctx: context.Background(),
				req: NewTbkItemInfoGetRequest("584399378309", "", 1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTopClient(tt.fields.appKey, tt.fields.appSecret)

			got, err := c.TbkItemInfoGet(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopClient.TbkItemInfoGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopClient.TbkItemInfoGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
