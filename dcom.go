package dcom

import (
	"github.com/DavidHsaiou/dcom/util"
)

type DCom interface {
	AddService(service any)
	Execute(execution any)
	BlockingRun()
	Run()
}

type Options interface {
	apply(d *dcomOption)
}

type dcom struct {
	container util.DiContainer
}

type dcomOption struct {
	ctx util.Context
}

func WithContext(ctx util.Context) Options {
	return withContextOption{
		ctx: ctx,
	}
}

type withContextOption struct {
	ctx util.Context
}

func (context withContextOption) apply(d *dcomOption) {
	d.ctx = context.ctx
}

func NewDCom(opt ...Options) DCom {
	var dcomOption dcomOption
	for _, o := range opt {
		o.apply(&dcomOption)
	}
	if dcomOption.ctx == nil {
		dcomOption.ctx = util.NewContext()
	}

	container := util.NewDIContainer(dcomOption.ctx)
	container.AddInstance(dcomOption.ctx)

	return &dcom{
		container: container,
	}
}

func (d *dcom) AddService(_ any) {
}

func (d *dcom) Execute(execution any) {
	d.container.AddExecute(execution)
}

func (d *dcom) BlockingRun() {
	d.container.Run()
}

func (d *dcom) Run() {
	go d.BlockingRun()
}
