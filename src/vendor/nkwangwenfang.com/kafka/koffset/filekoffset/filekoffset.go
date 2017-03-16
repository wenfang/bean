package filekoffset

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"nkwangwenfang.com/kafka/koffset"
)

const (
	poFile string = "%s/%s_partition_offset.%d"
)

type fileKOffset struct {
	app string
	dir string
}

// New 创建文件KOffsetter
func New(app, dir string) koffset.KOffsetter {
	return &fileKOffset{app: app, dir: dir}
}

func (fko *fileKOffset) Get(partitionID int32) (int64, error) {
	file := fmt.Sprintf(poFile, fko.dir, fko.app, partitionID)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	offset, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return 0, err
	}
	return offset, nil
}

func (fko *fileKOffset) Set(partitionID int32, offset int64) error {
	file := fmt.Sprintf(poFile, fko.dir, fko.app, partitionID)
	return ioutil.WriteFile(file, []byte(strconv.FormatInt(offset, 10)), 0666)
}
