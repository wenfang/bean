package splits

import (
	"testing"
)

func TestSplitUint64s(t *testing.T) {
	items := []uint64{1, 2, 3}
	t.Log(SplitUint64s(items, 1))
	t.Log(SplitUint64s(items, 2))
	t.Log(SplitUint64s(items, 3))
	t.Log(SplitUint64s(items, 4))
	t.Log(SplitUint64s(nil, 10))
}
