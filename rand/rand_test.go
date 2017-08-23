package rand

import (
	"testing"
)

func TestRandString(t *testing.T) {

	r1 := RandString(6)
	r2 := RandString(18)
	t.Log(r1)
	t.Log(r2)
	t.Log("=====")
	r3 := RandString(6, Lower)
	r4 := RandString(6, Upper)
	r5 := RandString(6, Number)
	t.Log(r3)
	t.Log(r4)
	t.Log(r5)

}

func TestRandStringMix(t *testing.T) {

	r1 := RandString(18, MixAll)
	r2 := RandString(18, MixLower)
	r3 := RandString(18, MixUpper)
	t.Log(r1)
	t.Log(r2)
	t.Log(r3)

}
func TestRandInt(t *testing.T) {
	t.Log(RandInt(16))
	t.Log(RandInt(8))

}

func BenchmarkRandString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(18)
	}

}

func BenchmarkRandInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(18)
	}

}
