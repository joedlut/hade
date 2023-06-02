package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter

	ctx context.Context

	hasTimeout bool
	//why use pointer ?
	writeMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writeMux:       &sync.Mutex{},
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writeMux
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		// func (u *URL) Query() Values
		//Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs.
		//To check errors use ParseQuery.

		//type Values map[string][]string
		//Values maps a string key to a list of values. It is typically used for query parameters and form values.
		//Unlike in the http.Header map, the keys in a Values map are case-sensitive.
		return map[string][]string(ctx.request.URL.Query())
	}
	return def
}

// implement the context.Context
// Done returns a channel that's closed when work done on behalf of this context should be canceled.

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}
