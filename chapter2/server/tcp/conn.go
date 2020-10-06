package tcp

import (
	"fmt"
	"net"
	"strconv"
)

// 读取 Key
func sendResponse(conn net.Conn, b []byte, err error) error {
	// 如果 err 不为空，则响应 -[错误内容长度] 错误内容
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d %s", len(errString), errString)
		_, err = conn.Write([]byte(tmp))
		return err
	}
	// 如果 err 为空，则响应 内容长度 内容
	_, err = conn.Write(append([]byte(strconv.Itoa(len(b))), b...))
	return err
}
