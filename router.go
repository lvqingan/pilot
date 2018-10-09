package pilot

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
	"strings"
)

type Router struct {

}

var (
	PRouter *Router
)

func init() {
	PRouter = &Router{}
}

// 初始化多路复用器
var router = httprouter.New()

// 处理静态文件
func (this *Router) HandleStatic(staticPath string, staticPrefix string) {
	staticHandler := http.FileServer(http.Dir(staticPath))
	http.Handle(staticPrefix, http.StripPrefix(staticPrefix, staticHandler))
}

// 处理get路由
func (this *Router) Get(pattern string, controller ControllerInterface, actionName string) {
	handler := func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		controllerName := reflect.ValueOf(controller).Type().String()
		controllerName = strings.TrimPrefix(controllerName, "*controllers.")
		controllerName = strings.TrimSuffix(controllerName, "Controller")

		controller.Init(strings.ToLower(controllerName), strings.ToLower(actionName))

		args := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
		reflect.ValueOf(controller).MethodByName(actionName).Call(args)
	}

	router.GET(pattern, handler)
}

// 获取多路复用器
func (this *Router) GetRouter() *httprouter.Router {
	return router
}
