package utils

import (
	"reflect"
	"testing"

	"github.com/tcchend/qcloud-svr/models"
)

func TestValidate(t *testing.T) {
	type args struct {
		tempItem *models.InstancesRequest
	}
	tests := []struct {
		name       string
		args       args
		wantNewResult *models.InstancesRequest
	}{
		{
			name:"validate",
			args:args{tempItem: &models.InstancesRequest{
				"aQYACCESSKEYIDEXAMPLE",
				"DescribeInstances",
				"2020-08-29T07:42:25Z",
				"20",
				"HmacSHA256",
				"1",
				"running",
				"2019-08-29T06:42:25Z",
				"1",
				"pek3a",
				"ihPnXFgsg5yyqhDN2IejJ2%2Bbo89ABQ1UqFkyOdzRITY%3D",
			}},
			wantNewResult: &models.InstancesRequest{
				"aQYACCESSKEYIDEXAMPLE",
				"DescribeInstances",
				"2020-08-29T07:42:25Z",
				"20",
				"HmacSHA256",
				"1",
				"running",
				"2019-08-29T06:42:25Z",
				"1",
				"pek3a",
				"ihPnXFgsg5yyqhDN2IejJ2%2Bbo89ABQ1UqFkyOdzRITY%3D",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewResult,_ := Validate(tt.args.tempItem); !reflect.DeepEqual(gotNewResult, tt.wantNewResult) {
				t.Errorf("Validate() = %v, want %v", gotNewResult, tt.wantNewResult)
			}
		})
	}
}
