package main

import (
	_ "GopherFocus/boot"
	_ "GopherFocus/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()

}
