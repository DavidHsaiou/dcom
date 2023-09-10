package util

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

var testServerUrl = "http://localhost:8080/abc"

func Test_httpClient_Get(t *testing.T) {
	defer gock.Off()
	gock.New(testServerUrl).
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test_httpClient_Get",
			args: args{
				url: testServerUrl,
			},
			wantErr: false,
		},
		{
			name: "Test_httpClient_Get_Error",
			args: args{
				url: testServerUrl + "/abc/123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHttpClient()
			resp, err := h.Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				DefaultLogger.Info(resp)
				require.NotEmpty(t, resp)
			}
		})
	}
}
