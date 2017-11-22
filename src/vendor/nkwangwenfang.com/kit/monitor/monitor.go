package monitor

import (
	"context"
	"net/http"

	"nkwangwenfang.com/log"
)

var monitorMux = http.NewServeMux()

// Handle 定义monitor对pattern的处理handler
func Handle(pattern string, handler http.Handler) { monitorMux.Handle(pattern, handler) }

// HandleFunc 定义monitor对patter的处理handler函数
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	monitorMux.HandleFunc(pattern, handler)
}

// Server Monitor Server结构
type Server struct {
	*http.Server
}

// New 创建一个新的Monitor Server
func New(config Config) *Server {
	var server Server
	if config.Addr != "" {
		server.Server = &http.Server{Addr: config.Addr, Handler: monitorMux}
		go func() {
			if err := server.ListenAndServe(); err != nil {
				log.Error("monitor server listen error", "addr", config.Addr, "error", err)
			}
		}()
	}
	return &server
}

func (server *Server) Shutdown(ctx context.Context) error {
	if server.Server == nil {
		return nil
	}
	return server.Server.Shutdown(ctx)
}

func (server *Server) Close() error {
	if server.Server == nil {
		return nil
	}
	return server.Server.Close()
}
