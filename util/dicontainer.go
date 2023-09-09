package util

import (
	gocontext "context"

	"go.uber.org/fx"
)

type DiContainer interface {
	AddLifetime(dependency LifetimeDependency)
	Add(dependency any)
	AddInstance(dependency any)
	AddExecute(executable any)
	Run()
	Stop()
}

type LifetimeDependency interface {
	OnStart(ctx gocontext.Context) error
	OnStop(ctx gocontext.Context) error
}

type diContainer struct {
	context              Context
	dependencies         []any
	lifetimeDependencies []LifetimeDependency
	executables          []any
	dependencyInstances  []any
	fxApp                *fx.App
}

func NewDIContainer(context Context) DiContainer {
	dependency := make([]any, 0)
	LifetimeDependencies := make([]LifetimeDependency, 0)
	dependencyInstances := make([]any, 0)
	return &diContainer{
		context:              context,
		dependencies:         dependency,
		lifetimeDependencies: LifetimeDependencies,
		dependencyInstances:  dependencyInstances,
	}
}

func (d *diContainer) AddLifetime(dependency LifetimeDependency) {
	d.lifetimeDependencies = append(d.lifetimeDependencies, dependency)
}

func (d *diContainer) Add(dependency any) {
	d.dependencies = append(d.dependencies, dependency)
}

func (d *diContainer) AddExecute(executable any) {
	d.executables = append(d.executables, executable)
}

func (d *diContainer) AddInstance(dependency any) {
	d.dependencyInstances = append(d.dependencyInstances, dependency)
}

// Run starts the DI container and blocks until the application is terminated.
func (d *diContainer) Run() {
	fxOptions := make([]fx.Option, 0)

	for _, dependency := range d.lifetimeDependencies {
		fxOptions = append(fxOptions, fx.Provide(withFxLifeCycle(dependency)))
	}

	for _, dependency := range d.dependencies {
		fxOptions = append(fxOptions, fx.Provide(dependency))
	}

	for _, execute := range d.executables {
		fxOptions = append(fxOptions, fx.Invoke(execute))
	}

	for _, inst := range d.dependencyInstances {
		fxOptions = append(fxOptions, fx.Supply(inst))
	}

	d.fxApp = fx.New(fxOptions...)
	go d.fxApp.Run()
	d.waitCancel()
}

func (d *diContainer) Stop() {
	d.fxApp.Stop(gocontext.Background())
}

func (d *diContainer) waitCancel() {
	<-d.context.Done()
	d.fxApp.Stop(gocontext.Background())
}

func withFxLifeCycle(dependency LifetimeDependency) func(lc fx.Lifecycle) any {
	return func(lc fx.Lifecycle) any {
		lc.Append(fx.Hook{
			OnStart: dependency.OnStart,
			OnStop:  dependency.OnStop,
		})
		return dependency
	}
}
