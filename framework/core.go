package framework

import (
	"net/http"
)

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

//实现Handler接口

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO
	panic("implement me")
}
