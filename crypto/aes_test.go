package crypto

import (
	"testing"
)

func TestAes(t *testing.T) {
	var source = "AA398!cs#"

	e, _ := Encrypt(source)
	t.Log(e)
	r, _ := Decrypt(e)
	t.Log(r)

	if r != source {
		t.Error("AES encrypt / decrypt error")
	}

}

func TestAesCoder(t *testing.T) {
	var v = "AA398!cs#"

	aesCoder := NewAesCoder()
	e, err := aesCoder.Encrypt(v)
	if err != nil {
		panic(err)
	}

	d, err := aesCoder.Decrypt(e)
	if err != nil {
		panic(err)
	}

	t.Log(e)
	t.Log(d)

}

func BenchmarkRandInt(b *testing.B) {
	var v = "LA791!cs#"
	for n := 0; n < b.N; n++ {
		Encrypt(v)
	}

}
