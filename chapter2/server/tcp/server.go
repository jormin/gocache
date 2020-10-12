package tcp

import (
	"bufio"
	"github.com/jormin/go-cache/chapter1/server/cache"
	"github.com/jormin/go-cache/helper"
	"io"
	"log"
	"net"
)

type Server struct {
	cache.Cache
}

// 监听 Tcp 服务
func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":9091")
	helper.Must(err)
	for {
		c, err := l.Accept()
		helper.Must(err)
		// 开启 goroutine
		go s.process(c)
	}
}

// 设置
func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	k, v, err := readKeyAndValue(r)
	if err != nil {
		return err
	}
	return sendResponse(conn, nil, s.Set(k, v))
}

// 读取
func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	k, err := readKey(r)
	if err != nil {
		return err
	}
	b, err := s.Get(k)
	return sendResponse(conn, b, err)
}

// 删除
func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	k, err := readKey(r)
	if err != nil {
		return err
	}
	return sendResponse(conn, nil, s.Delete(k))
}

// 处理
func (s *Server) process(conn net.Conn) {
	// 函数执行完毕关闭连接
	defer conn.Close()
	// Reader
	r := bufio.NewReader(conn)
	for {
		// 读取第一个字节
		op, err := r.ReadByte()
		if err != nil && err != io.EOF {
			log.Println("close connection due to error:", err)
			return
		}
		switch op {
		case 'S':
			// 设置
			err = s.set(conn, r)
		case 'G':
			// 读取
			err = s.get(conn, r)
		case 'D':
			// 删除
			err = s.del(conn, r)
		default:
			log.Println("close connection due to invalid operation", op)
			return
		}
		if err != nil {
			log.Println("close connection due to error:", err)
			return
		}
	}
}

// 生成服务
func New(c cache.Cache) *Server {
	return &Server{c}
}
