package main

import (
	"hade/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Addr: ":8080",
		//Handler: framework.NewCore()
		Handler: core,
	}
	server.ListenAndServe()
}
