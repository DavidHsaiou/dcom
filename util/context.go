package util

import (
	gocontext "context"
)

type Context interface {
	Cancel()
	Done() <-chan struct{}
}

type context struct {
	ctx    gocontext.Context
	cancel gocontext.CancelFunc
}

func NewContext() Context {
	ctx, cancel := gocontext.WithCancel(gocontext.Background())
	return &context{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (c *context) Cancel() {
	c.cancel()
}

func (c *context) Done() <-chan struct{} {
	return c.ctx.Done()
}
