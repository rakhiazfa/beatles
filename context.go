package beatles

import (
	"fmt"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

type defaultContext struct {
	app        Application
	requestCtx *fasthttp.RequestCtx

	request  Request
	response Response

	route        *route
	handlerIndex int
}

func NewContext(app Application) Context {
	return &defaultContext{
		app:          app,
		handlerIndex: -1,
	}
}

func (c *defaultContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (c *defaultContext) Done() <-chan struct{} {
	return nil
}

func (c *defaultContext) Err() error {
	return nil
}

func (c *defaultContext) Value(key any) any {
	return c.requestCtx.Value(key)
}

func (c *defaultContext) Application() Application {
	return c.app
}

func (c *defaultContext) RequestContext() *fasthttp.RequestCtx {
	return c.requestCtx
}

func (c *defaultContext) Request() Request {
	return c.request
}

func (c *defaultContext) Response() Response {
	return c.response
}

func (c *defaultContext) setRoute(route *route) {
	c.route = route
}

func (c *defaultContext) handlers() []Handler {
	if c.route == nil {
		return []Handler{}
	}

	return c.route.handlers
}

func (c *defaultContext) parameters() map[string]string {
	if c.route == nil {
		return make(map[string]string)
	}

	return c.route.parameters
}

func (c *defaultContext) Next() error {
	handlers := c.handlers()

	if len(handlers) == 0 {
		return NewHTTPError(http.StatusNotFound, fmt.Sprintf("Route %s not found", c.Request().Path()))
	}

	c.handlerIndex++

	if c.handlerIndex >= len(handlers) {
		return nil
	}

	return handlers[c.handlerIndex](c)
}

func (c *defaultContext) Reset(requestCtx *fasthttp.RequestCtx) {
	c.requestCtx = requestCtx

	c.request = newRequest(c)
	c.response = newResponse(c)
}

func (c *defaultContext) release() {
	c.requestCtx = nil

	c.request = nil
	c.response = nil

	c.route = nil
	c.handlerIndex = -1
}
