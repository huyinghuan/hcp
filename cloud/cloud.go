package cloud

type Cloud interface {
	//基本校验
	Verity()(error)
	// Put 文件上传
	Put() (string, error)
}
