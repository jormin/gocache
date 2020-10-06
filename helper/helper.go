package helper

// 异常处理
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
