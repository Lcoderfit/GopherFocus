package app

import "github.com/gogf/gf/frame/g"

// 启动应用
func Run() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// 静态目录设置
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		g.Log().Fatal("文件上传配置路径不能为空")
	}
}

/*

// 应用启动
func Run() {
	// 绑定Swagger Plugin
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// 静态目录设置
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		g.Log().Fatal("文件上传配置路径不能为空")
	}
	s.AddStaticPath("/upload", uploadPath)

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	// 业务系统初始化
	admin.Init()
	index.Init()

	// 启动Http Server
	s.Run()
}

*/
