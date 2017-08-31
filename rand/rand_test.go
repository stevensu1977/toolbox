package rand

import (
	"testing"
)

func TestRandStringNormal(t *testing.T) {
	result := []string{}

	result = append(result, RandString(6))
	result = append(result, RandString(18))

	result = append(result, RandString(18, Lower))
	result = append(result, RandString(18, Lower))
	result = append(result, RandString(18, Number))
	result = append(result, RandString(18, Hex))
	for index := range result {
		t.Log(result[index])
	}
}

func TestRandStringMix(t *testing.T) {
	result := []string{}

	result = append(result, RandString(18, MixAll))
	result = append(result, RandString(18, MixLower))
	result = append(result, RandString(18, MixUpper))
	for index := range result {
		t.Log(result[index])
	}

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
