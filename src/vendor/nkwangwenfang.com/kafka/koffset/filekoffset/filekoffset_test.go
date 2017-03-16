package filekoffset

import (
	"testing"
)

func TestFileOffset(t *testing.T) {
	app, dir := "testApp", "testdata"
	fo := New(app, dir)
	tryCase(t, fo, testCase{partition: 21, offset: 0})
	tryCase(t, fo, testCase{partition: 21, offset: 123})
}
