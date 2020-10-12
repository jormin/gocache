package http

import (
	"github.com/jormin/go-cache/chapter1/server/cache"
	"github.com/jormin/go-cache/helper"
	"net/http"
)

type Server struct {
	cache.Cache
}

// 监听 Http 服务
func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status/", s.statusHandler())
	err := http.ListenAndServe(":9090", nil)
	helper.Must(err)
}

// 返回缓存处理器
func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

// 返回状态处理器
func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}

// 生成服务
func New(c cache.Cache) *Server {
	return &Server{c}
}
