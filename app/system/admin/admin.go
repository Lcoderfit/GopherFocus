package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	g.Server().Group("/admin", func(group *ghttp.RouterGroup) {
		// 暂未开放
	})
}
