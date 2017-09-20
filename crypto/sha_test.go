package crypto

import (
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	sha := Sha1([]byte("Helloworld"))
	fmt.Println(sha)
}

func TestSha1File(t *testing.T) {
	sha, err := Sha1File("hello.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sha)
}
