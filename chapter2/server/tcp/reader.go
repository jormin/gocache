package tcp

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// 读取 Key
func readKey(r *bufio.Reader) (string, error) {
	l, err := readLen(r)
	if err != nil {
		return "", err
	}
	k := make([]byte, l)
	_, err = io.ReadFull(r, k)
	if err != nil {
		return "", err
	}
	return string(k), nil
}

// 读取 Key 和 Value
func readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	kl, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	vl, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	k := make([]byte, kl)
	_, err = io.ReadFull(r, k)
	if err != nil {
		return "", nil, err
	}
	v := make([]byte, vl)
	_, err = io.ReadFull(r, v)
	if err != nil {
		return "", nil, err
	}
	return string(k), v, err
}

// 读取长度
func readLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return l, nil
}
