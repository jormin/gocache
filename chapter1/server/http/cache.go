package http

import (
	"encoding/json"
	"github.com/jormin/go-cache/helper"
	"github.com/jormin/go-cache/response"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 缓存处理器
type cacheHandler struct {
	*Server
}

// 处理Http服务
func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取Key
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 读取不到Key，响应400
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	// 请求方法
	m := r.Method
	switch m {
	// 读取
	case http.MethodGet:
		b, err := h.Get(key)
		// 读取异常则记录异常并响应500
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 如果读取内容长度为0，则响应404
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		// 正常读取则响应缓存内容
		b, _ = json.Marshal(response.Response{
			Code:    0,
			Message: "get cache success",
			Data:    string(b),
		})
		_, _ = w.Write(b)
	// 设置
	case http.MethodPut:
		// 读取值
		b, err := ioutil.ReadAll(r.Body)
		helper.Must(err)
		if len(b) != 0 {
			err := h.Set(key, b)
			// 设置错误则记录异常并响应500
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		b, _ = json.Marshal(response.Response{
			Code:    0,
			Message: "put cache success",
		})
		_, _ = w.Write(b)
	// 删除
	case http.MethodDelete:
		err := h.Delete(key)
		// 删除错误则记录异常并响应500
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		b, _ := json.Marshal(response.Response{
			Code:    0,
			Message: "delete cache success",
		})
		_, _ = w.Write(b)
	}
	// 正常响应200
	w.WriteHeader(http.StatusOK)
}
