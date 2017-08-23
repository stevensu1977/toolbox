package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/stevensu1977/toolbox/rand"
)

var (
	KEY_HEADER = "AES256Key-"
	KEY        = []byte("AES256Key-32Characters1234567890")
	HEX_STR    = "37b8e8a308c354048d245f6d"
)

type AESCoder struct {
	Hex string
	Key []byte
}

func (aesCoder *AESCoder) Encrypt(plain string) (string, error) {
	return encrypt(plain, aesCoder.Key, aesCoder.Hex)
}

func (aesCoder *AESCoder) Decrypt(plain string) (string, error) {
	return decrypt(plain, aesCoder.Key, aesCoder.Hex)
}

func NewAesCoder() *AESCoder {

	return &AESCoder{
		Key: []byte(KEY_HEADER + rand.RandString(22)),
		Hex: rand.RandString(24, rand.Hex),
	}
}

func Encrypt(plain string) (string, error) {

	return encrypt(plain, KEY, HEX_STR)
}

func encrypt(plain string, key []byte, hex_str string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce, _ := hex.DecodeString(hex_str)
	//fmt.Println(nonce)
	//nonce := []byte{55, 184, 232, 163, 8, 195, 84, 4, 141, 36, 95, 109}
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plain), nil)
	return hex.EncodeToString(ciphertext[:]), nil
}

func Decrypt(plain string) (string, error) {

	return decrypt(plain, KEY, HEX_STR)
}

func decrypt(plain string, key []byte, hex_str string) (string, error) {

	nonce, _ := hex.DecodeString(hex_str)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	hexPlain, err := hex.DecodeString(plain)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, []byte(hexPlain[:]), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext[:]), nil
}
