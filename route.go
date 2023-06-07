package main

import "hade/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
