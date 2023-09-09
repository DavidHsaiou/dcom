package dcom

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DavidHsaiou/dcom/util"
)

func TestNewDCom(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "NewDCom",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			com := NewDCom()
			if tt.wantErr {
				assert.Nil(t, com)
			} else {
				assert.NotNil(t, com)
			}
		})
	}
}

func Test_dcom_AddService(t *testing.T) {
	type fields struct {
		container util.DiContainer
	}
	type args struct {
		service any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "AddService",
			args: args{
				service: func() {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dcom{
				container: tt.fields.container,
			}
			d.AddService(tt.args.service)
		})
	}
}

func Test_dcom_Run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "BlockingRun",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := util.NewContext()
			d := NewDCom(WithContext(ctx))
			d.Execute(func() { print("hello") })
			d.Run()
			time.Sleep(1 * time.Second)
			ctx.Cancel()
		})
	}
}
