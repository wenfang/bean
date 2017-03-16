package conv

import (
	"testing"
)

var tableStringInt = map[string]int{
	"1234234": 1234234,
	"1123":    1123,
	"010":     10,
}

var tableStringInt64 = map[string]int64{
	"100234234": 100234234,
	"2349000":   2349000,
	"0324":      324,
}

var tableFloat = map[string]float64{
	"1.0":         1,
	"10.34":       10.34,
	"1341.000000": 1341,
}

var tableString = map[float64]string{
	10.34:    "10.34",
	1.0000:   "1",
	12341234: "12341234",
}

func TestConvInt(t *testing.T) {
	for s, num := range tableStringInt {
		i, err := ItoInt(s)
		if err != nil {
			t.Fatal(err)
		}
		if i != num {
			t.Fatal("not equal", i, num)
		}
	}
}

func TestConvInt64(t *testing.T) {
	for s, num := range tableStringInt64 {
		i, err := ItoInt64(s)
		if err != nil {
			t.Fatal(err)
		}
		if i != num {
			t.Fatal("not equal", i, num)
		}
	}
}

func TestConvFloat64(t *testing.T) {
	for s, num := range tableFloat {
		i, err := ItoFloat64(s)
		if err != nil {
			t.Fatal(err)
		}
		if i != num {
			t.Fatal("not equal", i, num)
		}
	}
}

func TestConvString(t *testing.T) {
	for s, num := range tableString {
		i, err := ItoString(s)
		if err != nil {
			t.Fatal(err)
		}
		if i != num {
			t.Fatal("not equal", i, num)
		}
	}
}

func TestUint(t *testing.T) {
	a := uint(10)
	i, err := ItoInt(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(i)
}
