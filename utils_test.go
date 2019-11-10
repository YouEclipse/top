package top

import (
	"testing"
)

func Test_kvPairList_load(t *testing.T) {
	type fields struct {
		list []*kvPair
	}

	type A struct {
		B string `json:"b"`
		C bool   `json:"c"`
		D int64  `json:"d"`
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				list: make([]*kvPair, 0),
			},
			args: args{
				data: A{"你好呀", false, 666},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &kvPairList{
				list: tt.fields.list,
			}
			if err := l.load(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("kvPairList.load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
