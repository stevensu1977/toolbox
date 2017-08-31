package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"

	"github.com/stevensu1977/toolbox/rand"
)

var (
	KEY_HEADER = "AES256Key-"
	KEY        = []byte("AES256Key-32Characters1234567890")
	HEX_STR    = "37b8e8a308c354048d245f6d"
	AES_Prefix = "{AES}"
)

// AESCoder is abstract struct ,provid all AES encrypt, decrypt func
type AESCoder struct {
	Hex string
	Key []byte
}

func haveAesPrefix(payload string) bool {
	return strings.HasPrefix(AES_Prefix)
}

//Encrypt provide AES encrypt , accept string
func (aesCoder *AESCoder) Encrypt(plain string) (string, error) {
	return aesEncrypt(plain, aesCoder.Key, aesCoder.Hex)
}

//Decrypt provide AES decrypt , accept string
func (aesCoder *AESCoder) Decrypt(plain string) (string, error) {
	return aesDecrypt(plain, aesCoder.Key, aesCoder.Hex)
}

//NewAesCoder get
func NewAesCoder(key ...string) *AESCoder {
	aesKey := []byte(KEY_HEADER + rand.RandString(22))
	aesHex := rand.RandString(24, rand.Hex)

	if len(key) == 2 && len(key[0]) == 32 && len(key[1]) == 24 {
		aesKey = []byte(key[0])
		aesHex = key[1]
	}

	return &AESCoder{
		Key: []byte(aesKey),
		Hex: aesHex,
	}
}

func AESEncrypt(plain string) (string, error) {

	return aesEncrypt(plain, KEY, HEX_STR)
}

func aesEncrypt(plain string, key []byte, hex_str string) (string, error) {
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

func AESDecrypt(plain string) (string, error) {

	return aesDecrypt(plain, KEY, HEX_STR)
}

func aesDecrypt(plain string, key []byte, hex_str string) (string, error) {

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
