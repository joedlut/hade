package framework

import (
	"fmt"
	"log"
	"net/http"
)

type Core struct {
	//控制器
	router map[string]ControllerHandler
}

func NewCore() *Core {
	//return &Core{}
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

//实现Handler接口

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//TODO
	//panic("implement me")
	log.Println("core.serveHttp")
	ctx := NewContext(request, response)

	router := c.router["foo"]
	//if router != nil {
	if router == nil {
		return
	}
	log.Println("core.router")

	// ???
	// 执行控制器的函数
	fmt.Printf("%T", router)
	fmt.Println("------")
	fmt.Println(ctx.request.URL)
	fmt.Println(ctx.request.Body)
	fmt.Println(ctx.request.Header)
	fmt.Println(ctx.request.RemoteAddr)
	fmt.Println(ctx.request.Host)

	router(ctx)
}
