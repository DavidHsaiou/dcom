package server

import (
	gocontext "context"
	gohttp "net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/DavidHsaiou/dcom/dto"
	"github.com/DavidHsaiou/dcom/util"
)

var testServerAddr = "localhost:8081"
var testServerUrl = "http://" + testServerAddr

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
					WithAddr(testServerAddr),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewHTTP(tt.args.opts...)
			server.AddRoute(new(TestRoute))
			go server.Run()
			time.Sleep(1 * time.Second)
			resp, err := util.NewHttpClient().Get(testServerUrl)
			if err != nil && !tt.wantErr {
				require.NoError(t, err)
			}
			util.DefaultLogger.Info(resp)
			assert.NotEmptyf(t, resp, "resp should not be empty")
			server.Stop()
		})
	}
}

func TestStartAndStop(t *testing.T) {
	type args struct {
		opts []Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestStartAndStop",
			args: args{
				opts: []Options{
					WithAddr(testServerAddr),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewHTTP(tt.args.opts...)
			server.AddRoute(TestRoute{})
			err := server.OnStart(gocontext.Background())
			if err != nil && !tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			time.Sleep(1 * time.Second)
			err = server.OnStop(gocontext.Background())
			if err != nil && !tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
