package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// 状态处理器
type statusHandler struct {
	*Server
}

// 处理Http服务
func (h statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 查询状态只支持Get请求，非Get请求一律响应405
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// 获取并json打包数据
	b, err := json.Marshal(h.GetStat())
	// json打包失败则记录异常并响应500
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(b)
}
