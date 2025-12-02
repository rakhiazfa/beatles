package beatles

import (
	"errors"
	"net"
	"sync"

	"github.com/rakhiazfa/beatles/logger"
	"github.com/valyala/fasthttp"
)

type application struct {
	pool sync.Pool

	logger logger.Logger

	config ApplicationConfig
	server *fasthttp.Server

	router RouterTree
}

func New(config ...ApplicationConfig) Application {
	app := &application{
		logger: logger.New(),
		config: mergeApplicationConfig(DefaultApplicationConfig, config...),
		router: NewRouterTree(),
	}

	app.pool = sync.Pool{
		New: func() any {
			return NewContext(app)
		},
	}

	app.server = &fasthttp.Server{
		LogAllErrors:                 false,
		Handler:                      app.handler(),
		Concurrency:                  app.config.Concurrency,
		ReadBufferSize:               app.config.ReadBufferSize,
		WriteBufferSize:              app.config.WriteBufferSize,
		ReadTimeout:                  app.config.ReadTimeout,
		WriteTimeout:                 app.config.WriteTimeout,
		IdleTimeout:                  app.config.IdleTimeout,
		MaxRequestBodySize:           app.config.MaxRequestBodySize,
		DisableKeepalive:             app.config.DisableKeepalive,
		ReduceMemoryUsage:            app.config.ReduceMemoryUsage,
		DisablePreParseMultipartForm: app.config.DisablePreParseMultipartForm,
		StreamRequestBody:            app.config.StreamRequestBody,
	}

	return app
}

func (app *application) Logger() logger.Logger {
	return app.logger
}

func (app *application) Server() *fasthttp.Server {
	return app.server
}

func (app *application) Router() RouterTree {
	return app.router
}

func (app *application) AcquireContext(requestCtx *fasthttp.RequestCtx) Context {
	c, ok := app.pool.Get().(Context)
	if !ok {
		panic(errors.New("sync.Pool returned value with invalid type: expected Context"))
	}

	c.Reset(requestCtx)

	return c
}

func (app *application) ReleaseContext(c Context) {
	c.release()

	app.pool.Put(c)
}

func (app *application) handler() fasthttp.RequestHandler {
	return func(requestCtx *fasthttp.RequestCtx) {
		c := app.AcquireContext(requestCtx)
		defer app.ReleaseContext(c)

		route, err := app.router.Search(c.Request().Method(), c.Request().Path())
		if err != nil {
			// TODO: handle error
			return
		}

		c.setRoute(route)

		if err := c.Next(); err != nil {
			// TODO: handle error
			return
		}
	}
}

func (app *application) Start(addr string, config ...StartConfig) error {
	mergedConfig := mergeStartConfig(DefaultStartConfig, config...)

	listener, err := net.Listen(mergedConfig.Network, addr)
	if err != nil {
		return err
	}

	app.logger.Infof("Server started on %s", listener.Addr().String())

	if mergedConfig.CertFile != "" && mergedConfig.KeyFile != "" {
		return app.server.ServeTLS(listener, mergedConfig.CertFile, mergedConfig.KeyFile)
	}

	return app.server.Serve(listener)
}

func (app *application) Stop() error {
	return nil
}
