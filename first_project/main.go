package main

import (
	_ "first_project/boot"
	_ "first_project/router"
	"github.com/gogf/gf/g"
)

func main() {
	_ = g.Server().Run()
}
