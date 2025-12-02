package beatles

import (
	"context"

	"github.com/valyala/fasthttp"
)

type Context interface {
	context.Context

	// Returns the application instance
	Application() Application

	// Returns *fasthttp.RequestCtx that carries a deadline
	RequestContext() *fasthttp.RequestCtx

	// Returns the request instance
	Request() Request

	// Returns the response instance
	Response() Response

	// Stores the given route into the context
	setRoute(route *Route)

	// Returns the handlers extracted from the matched route
	Handlers() []Handler

	// Returns the route parameters extracted from the matched route
	Parameters() map[string]string

	// Executes the next method in the handler stack that matches the current route
	Next() error

	// Reset context fields based on the request from the handler
	Reset(requestCtx *fasthttp.RequestCtx)

	// Resets all internal fields of Context so the instance can be safely reused by the context pool
	release()
}
