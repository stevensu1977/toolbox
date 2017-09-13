package crypto

import (
	"encoding/base64"
	"strings"
)

//BasicAuthEncode is HTTP Basic Auth helper func
func BasicAuthEncode(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//BasicAuthDecode is HTTP Basic Auth helper func
func BasicAuthDecode(input string) ([]string, error) {

	data, err := Base64Decode(input)

	if err != nil {
		return nil, err
	}

	return strings.Split(data, ":"), nil

}

//Base64Encode simple base64 helper func
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

//Base64Decode simple base64 helpler func
func Base64Decode(input string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
