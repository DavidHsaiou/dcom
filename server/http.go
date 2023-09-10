package server

import (
	gocontext "context"
	"errors"
	gohttp "net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/DavidHsaiou/dcom/dto"
	"github.com/DavidHsaiou/dcom/util"
)

type Route interface {
	Method() string
	Path() string
	Handler(request *dto.Request) *dto.Response
	Group() string
}

type Handler func(request *dto.Request) *dto.Response

type HTTP interface {
	AddRoute(route Route)
	Run()
	Stop()

	util.LifetimeDependency
}

type Options interface {
	apply(options *options)
}

type options struct {
	addr              string
	readHeaderTimeout time.Duration
}

var (
	defaultOptions = &options{
		addr:              ":8080",
		readHeaderTimeout: 10 * time.Second,
	}
)

type http struct {
	engine *gin.Engine
	server *gohttp.Server
}

func NewHTTP(opts ...Options) HTTP {
	for _, opt := range opts {
		opt.apply(defaultOptions)
	}

	engine := gin.Default()
	return &http{
		engine: engine,
		server: &gohttp.Server{
			Addr:              defaultOptions.addr,
			Handler:           engine,
			ReadHeaderTimeout: defaultOptions.readHeaderTimeout,
		},
	}
}

func (h *http) AddRoute(route Route) {
	var routeEngine any
	if route.Group() != "" {
		routeEngine = h.engine.Group(route.Group())
	} else {
		routeEngine = h.engine
	}

	params := make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf(route.Path()))
	params = append(params, reflect.ValueOf(decorateWithGinHandler(route.Handler)))

	reflect.
		ValueOf(routeEngine).
		MethodByName(route.Method()).
		Call(params)
}

func (h *http) Run() {
	if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, gohttp.ErrServerClosed) {
		panic(err)
	}
}

func (h *http) Stop() {
	if err := h.server.Shutdown(gocontext.Background()); err != nil {
		panic(err)
	}
}

func (h *http) OnStart(_ gocontext.Context) error {
	go h.Run()
	return nil
}

func (h *http) OnStop(_ gocontext.Context) error {
	h.Stop()
	return nil
}

func WithAddr(addr string) Options {
	return withAddrOption(addr)
}

func decorateWithGinHandler(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &dto.Request{
			Data: c,
		}
		response := handler(request)

		c.JSON(200, response)
	}
}

type withAddrOption string

func (o withAddrOption) apply(options *options) {
	options.addr = string(o)
}
