package utils

import (
	"testing"
)

func TestRand(t *testing.T) {

	r1 := RandString(10)
	r2 := Rand(10, Number)
	t.Log(r1)
	t.Log(r2)
}
