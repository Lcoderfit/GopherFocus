package service

import (
	"GopherFocus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 试图管理服务
var View = viewService{}

type viewService struct{}

// 渲染默认模板页面
func (s *viewService) Render(r *ghttp.Request, data ...model.View) {
	s.RenderTp
}

// 渲染指定模板页面
func (s *viewService) RenderTpl(r *ghttp.Request, tpl string, data ...model.View) {
	var (
		viewObj  = model.View{}
		ViewData = make(g.Map)
	)

	if len(data) > 0 {
		viewObj = data[0]
	}

	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().GetString("setting.title")
	} else {
		viewObj.Title = viewObj.Title + `-` + g.Cfg().GetString(``)
	}
}

/*
// 渲染指定模板页面
func (s *viewService) RenderTpl(r *ghttp.Request, tpl string, data ...model.View) {
	var (
		viewObj  = model.View{}
		viewData = make(g.Map)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().GetString(`setting.title`)
	} else {
		viewObj.Title = viewObj.Title + ` - ` + g.Cfg().GetString(`setting.title`)
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = g.Cfg().GetString(`setting.keywords`)
	}
	if viewObj.Description == "" {
		viewObj.Description = g.Cfg().GetString(`setting.description`)
	}
	// 去掉空数据
	viewData = gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}
	// 内置对象
	viewData["BuildIn"] = &viewBuildIn{httpRequest: r}
	// 内容模板
	if viewData["MainTpl"] == nil {
		viewData["MainTpl"] = s.getDefaultMainTpl(r)
	}
	// 提示信息
	if notice, _ := Session.GetNotice(r.Context()); notice != nil {
		_ = Session.RemoveNotice(r.Context())
		viewData["Notice"] = notice
	}
	// 渲染模板
	_ = r.Response.WriteTpl(tpl, viewData)
	// 开发模式下，在页面最下面打印所有的模板变量
	if gmode.IsDevelop() {
		_ = r.Response.WriteTplContent(`{{dump .}}`, viewData)
	}
	// 退出当前业务函数执行
	r.Exit()
}


// 渲染默认模板页面
func (s *viewService) Render(r *ghttp.Request, data ...model.View) {
	s.RenderTpl(r, g.Cfg().GetString("viewer.indexLayout"), data...)
}

// 跳转中间页面
func (s *viewService) Render302(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "页面跳转中"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/302.html"
	s.Render(r, view)
}

// 401页面
func (s *viewService) Render401(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/401.html"
	s.Render(r, view)
}

// 403页面
func (s *viewService) Render403(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/403.html"
	s.Render(r, view)
}

// 404页面
func (s *viewService) Render404(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "资源不存在"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/404.html"
	s.Render(r, view)
}

// 500页面
func (s *viewService) Render500(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "请求执行错误"
	}
	view.MainTpl = s.getViewFolderName() + "/pages/500.html"
	s.Render(r, view)
}
*/
