package crypto

import (
	"testing"
)

func TestBasicAuthEncode(t *testing.T) {

	encodePassword := BasicAuthEncode("wei.su@oracle.com", "a123gjGJ")
	t.Log(encodePassword)
	raw, err := Base64Decode(Base64Encode("welcome1"))
	if err != nil {
		t.Fatal(err)
	}
	if raw != "welcome1" {
		t.Fatal("not match")
	}

	//str1, err := Base64Decode("dx1212313=")
	str1, err := Base64Decode(encodePassword)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str1)
	t.Log(BasicAuthDecode(encodePassword))

}
