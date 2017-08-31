package rand

import (
	"testing"
)

func TestRandStringNormal(t *testing.T) {
	result := []string{}

	result = append(result, RandString(6))
	result = append(result, RandString(18))

	result = append(result, RandString(6, Lower))
	result = append(result, RandString(6, Lower))
	result = append(result, RandString(6, Number))

	for index := range result {
		t.Log(result[index])
	}
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
