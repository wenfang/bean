package koffset

// KOffsetter 接口
type KOffsetter interface {
	Get(partitionID int32) (int64, error)
	Set(partitionID int32, offset int64) error
}
