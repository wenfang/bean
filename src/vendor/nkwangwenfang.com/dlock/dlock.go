package dlock

// Locker 接口
type Locker interface {
	Lock() (bool, error)
	Unlock() (bool, error)
}
