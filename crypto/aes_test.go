package crypto

import (
	"testing"
)

func TestAes(t *testing.T) {
	var source = "AA398!cs#"
	e, _ := AESEncrypt(source)
	t.Log(e)
	r, _ := AESDecrypt(e)
	t.Log(r)

	if r != source {
		t.Error("AES encrypt / decrypt error")
	}

}
func TestAesPrefix(t *testing.T) {
	d, _ := AESEncrypt("AA398!cs#")

	var payload = "{AES}" + d
	if ok := HasAesPrefix(payload); ok {
		if e, err := AESDecrypt(RemoveAesPrefix(payload)); e != "AA398!cs#" || err != nil {
			t.Fatal("decrypt error")
		}
		t.Log("OK")
	}

}
func TestAesCoder(t *testing.T) {
	var payload = "AA398!cs#"

	aesCoder := NewAesCoder()
	eString, err := aesCoder.Encrypt(payload)
	if err != nil {
		panic(err)
	}
	dString, err := aesCoder.Decrypt(eString)
	if err != nil {
		panic(err)
	}

	if dString != payload {
		t.Fatal("encrypt, decrypt error!")
	}

	t.Logf("payload %s , encrypt %s , decrypt  %s", payload, eString, dString)

}

func BenchmarkRandInt(b *testing.B) {
	var payload = "LA791!cs#"
	for n := 0; n < b.N; n++ {
		AESEncrypt(payload)
	}

}
