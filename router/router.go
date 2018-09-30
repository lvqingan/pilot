package router

import (
	"net/http"
)

// 初始化多路复用器
var mux = http.NewServeMux()

// 处理静态文件
func HandleStatic(staticPath string, staticPrefix string)  {
	staticHandler := http.FileServer(http.Dir(staticPath))
	mux.Handle(staticPrefix, http.StripPrefix(staticPrefix, staticHandler))
}

// 处理路由
func Handle(callback func(mux *http.ServeMux)) {
	callback(mux)
}

// 获取多路复用器
func GetMux() *http.ServeMux {
	return mux
}
