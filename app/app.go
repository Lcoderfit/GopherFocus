package app

import (
	"GopherFocus/app/system/admin"
	"GopherFocus/app/system/index"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmode"
	"github.com/gogf/swagger"
)

// 启动应用
func Run() {
	//g.Log().Info("gmode: ", gmode.IsDevelop())
	//a := gcmd.GetOptWithEnv("gf.fuck")
	//g.Log().Info("gfuck: ", a)
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// 静态目录设置
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		g.Log().Fatal("文件上传配置路径不能为空")
	}
	// SetServerRoot: 配置默认静态文件查找路径
	// AddSearchPath: 添加静态文件查找路径
	// AddStaticPath: 配置URI与静态文件路径的映射关系， 深度优先匹配原则，url层级越深优先级越高
	s.AddStaticPath("/upload", uploadPath)

	// hook，开发阶段禁止浏览器缓存，便于调试
	if gmode.IsDevelop() {
		// hook函数
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			// no-store表示不缓存请求或响应的任何内容，这里设置在响应首部中，表示告知客户端不缓存响应内容
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	// 业务系统初始化
	admin.Init()
	index.Init()

	// 启动Http Server
	s.Run()
}
