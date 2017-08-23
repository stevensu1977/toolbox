package utils

import (
	"math/rand"
	"time"
)

const (
	Mix = iota
	Number
	Upper
	Lower
	MixUpper
	MixLower
)

func rand_seek(rand_type int) string {

	switch rand_type {
	case Number:
		return "0123456789"
	case Upper:
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case Lower:
		return "abcdefghijklmnopqrstuvwxyz"
	case MixUpper:
		return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case MixLower:
		return "0123456789abcdefghijklmnopqrstuvwxyz"
	default:
		return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
}

func RandString(str_size int) string {

	mix := Mix
	return Rand(str_size, mix)
}

func Rand(str_size int, rand_type int) string {

	alphanum := rand_seek(rand_type)
	return rand_str(str_size, alphanum)
}

func rand_str(str_size int, alphanum string) string {
	var bytes = make([]byte, str_size)
	rand.Seed(time.Now().UTC().UnixNano()) //very important
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
