package server

import (
	gohttp "net/http"
	"testing"
	"time"

	"github.com/DavidHsaiou/dcom/dto"
)

type TestRoute struct {
}

func (t TestRoute) Method() string {
	return gohttp.MethodGet
}

func (t TestRoute) Path() string {
	return "/"
}

func (t TestRoute) Handler(_ *dto.Request) *dto.Response {
	return dto.NewResponse(dto.ResultCodeSuccess, nil)
}

func (t TestRoute) Group() string {
	return ""
}

func TestNewHttp(t *testing.T) {
	type args struct {
		opts []Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNewHttp",
			args: args{
				opts: []Options{
					WithAddr(":8081"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewHTTP(tt.args.opts...)
			server.AddRoute(TestRoute{})
			go server.Run()
			time.Sleep(1 * time.Second)
			server.Stop()
		})
	}
}
