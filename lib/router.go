package lib

import (
	"net/http"
	"reflect"
	"strings"
)

type Router struct {

}

func GetRouter() *Router {
	router := &Router{}

	return router
}

// 初始化多路复用器
var mux = http.NewServeMux()

// 处理静态文件
func (this *Router) HandleStatic(staticPath string, staticPrefix string) {
	staticHandler := http.FileServer(http.Dir(staticPath))
	mux.Handle(staticPrefix, http.StripPrefix(staticPrefix, staticHandler))
}

// 处理get路由
func (this *Router) Get(pattern string, controller ControllerInterface, actionName string) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		controller.Init("home", strings.ToLower(actionName))

		args := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
		reflect.ValueOf(controller).MethodByName(actionName).Call(args)
	}

	mux.HandleFunc(pattern, handler)
}

// 获取多路复用器
func (this *Router) GetMux() *http.ServeMux {
	return mux
}
