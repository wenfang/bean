package monitor

import (
	"context"
	"errors"
	"net/http"
)

var (
	errMonitorRunning = errors.New("monitor server is running")
)

var (
	monitorMux    = http.NewServeMux()
	monitorServer *http.Server
)

// Handle 定义monitor对pattern的处理handler
func Handle(pattern string, handler http.Handler) { monitorMux.Handle(pattern, handler) }

// HandleFunc 定义monitor对patter的处理handler函数
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	monitorMux.HandleFunc(pattern, handler)
}

// ListenAndServe 监听tcp地址启动服务
func ListenAndServe(addr string) error {
	if monitorServer != nil {
		return errMonitorRunning
	}
	monitorServer = &http.Server{Addr: addr, Handler: monitorMux}
	return monitorServer.ListenAndServe()
}

// Close 关闭monitorServer
func Close() error {
	if monitorServer == nil {
		return nil
	}
	return monitorServer.Close()
}

// Shutdown 优雅关闭monitorServer
func Shutdown(ctx context.Context) error {
	if monitorServer == nil {
		return nil
	}
	return monitorServer.Shutdown(ctx)
}
