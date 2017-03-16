package rdskoffset

import (
	"testing"

	"nkwangwenfang.com/kafka/koffset"
)

type testCase struct {
	partition int32
	offset    int64
}

func tryCase(t *testing.T, off koffset.KOffsetter, test testCase) {
	if err := off.Set(test.partition, test.offset); err != nil {
		t.Fatal(err)
	}

	offset, err := off.Get(test.partition)
	if err != nil {
		t.Fatal(err)
	}

	if offset != test.offset {
		t.Fatal("offset not equal")
	}
}
