package util

import (
	gocontext "context"
	"testing"
	"time"
)

var logger = NewZapLogger()

func TestNewDIContainer(t *testing.T) {
	type args struct {
		context Context
	}
	ctx := NewContext()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestNewDIContainer",
			args: args{
				context: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDIContainer(tt.args.context)
			d.AddLifetime(LifetimeDependencyMock{})
			d.AddInstance(logger)
			d.AddExecute(func(log *zapLogger) {
				log.Info("test")
			})
			go d.Run()
			time.Sleep(1 * time.Second)
			tt.args.context.Cancel()
		})
	}
}

type LifetimeDependencyMock struct {
}

func (l LifetimeDependencyMock) OnStart(_ gocontext.Context) error {
	logger.Info("OnStart")
	return nil
}

func (l LifetimeDependencyMock) OnStop(_ gocontext.Context) error {
	logger.Info("OnStop")
	return nil
}
