package handler1

import (
	"reflect"
	"testing"
)

type Dump struct {
	I []int
}

func TestInput2Value(t *testing.T) {
	dump := &Dump{}
	sv := reflect.ValueOf(dump)
	sv = sv.Elem()

	err := input2Value("1,2,3", sv.FieldByName("I"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dump.I)
}
