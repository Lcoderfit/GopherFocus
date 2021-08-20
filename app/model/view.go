package model

// 视图渲染内容对象
type View struct {
	Title       string           // 页面标题
	Keywords    string           // 页面关键字
	Description string           // 页面描述
	Error       string           // 错误信息
	MainTpl     string           // 自定义MainTpl展示模板文件
	Redirect    string           // 引导页面跳转
	ContentType string           // 内容模型
	BreadCrumb  []ViewBreadCrumb // 面包屑
	Data        interface{}      // 页面参数v
}

// 视图面包屑结构
type ViewBreadCrumb struct {
	Name string // 显示名称
	Url  string // 链接地址，当为空时表示被选中
}
