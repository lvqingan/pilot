package lib

import (
	"html/template"
	"net/http"
)

type ControllerInterface interface {
	View(writer http.ResponseWriter, data interface{})
	Init(controllerName string, actionName string)
}

type Controller struct {
	ControllerName string
	ActionName     string
}

func (this *Controller) View(writer http.ResponseWriter, data interface{}) {
	// 1. 解析模板
	tplPath := GetConfig().GetString("app.template_path") + "/" + this.ControllerName + "/" + this.ActionName + ".html"
	tpl := template.Must(template.ParseFiles(tplPath))

	// 2. 渲染模板
	tpl.ExecuteTemplate(writer, this.ActionName, data)
}

func (this *Controller) Init(controllerName string, actionName string) {
	this.ControllerName = controllerName
	this.ActionName = actionName
}
