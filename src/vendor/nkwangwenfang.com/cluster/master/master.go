package master

// Master 接口，判断是否为Master
type Master interface {
	// CheckMaster 检查并获取master权限
	// 返回true获取成功为Master
	// 返回false,error为nil,获取成功,非Master
	// false, error非nil,获取失败
	CheckMaster() (bool, error)
}
