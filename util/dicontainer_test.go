package util

import (
	"testing"
	"time"
)

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
			go d.Run()
			time.Sleep(1 * time.Second)
			tt.args.context.Cancel()
		})
	}
}
