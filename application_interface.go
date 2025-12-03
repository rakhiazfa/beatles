package beatles

import (
	"github.com/rakhiazfa/beatles/logger"
	"github.com/valyala/fasthttp"
)

type Application interface {
	// Returns the logger instance
	Logger() logger.Logger

	// Returns the *fasthttp.Server
	Server() *fasthttp.Server

	// Returns the router tree instance
	Router() RouterTree

	// Stores the given error handler into the application
	SetErrorHandler(errorHandler ErrorHandler) Application

	// Retrieves a Context instance from the sync.Pool
	AcquireContext(requestCtx *fasthttp.RequestCtx) Context

	// Release Context instance back to sync.Pool
	ReleaseContext(c Context)

	// Returns the fasthttp.RequestHandler used to process incoming HTTP requests
	handler() fasthttp.RequestHandler

	// Starts the HTTP server on given address using the provided optional ListenConfig
	Start(addr string, config ...StartConfig) error
}
