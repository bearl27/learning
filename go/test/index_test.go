package test

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 2
	if Debug {
		t.Skip("skip")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
