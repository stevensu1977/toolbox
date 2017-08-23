package rand

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	Mix = iota
	Number
	NumberNoZero
	Hex
	Upper
	Lower
	MixAll
	MixUpper
	MixLower
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano()) //very important
}

func rand_seek(rand_type int) string {

	switch rand_type {
	case Number:
		return "0123456789"
	case NumberNoZero:
		return "123456789"
	case Hex:
		return "0123456789ABCDEF"
	case Upper:
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case Lower:
		return "abcdefghijklmnopqrstuvwxyz"
	case MixAll:
		return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case MixUpper:
		return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case MixLower:
		return "0123456789abcdefghijklmnopqrstuvwxyz"
	default:
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
}

func RandString(str_size int, rand_type ...int) string {

	mix := Mix
	if len(rand_type) > 0 {
		mix = rand_type[0]
	}
	return _rand(str_size, mix)
}

func _rand(str_size int, rand_type int) string {

	alphanum := rand_seek(rand_type)
	return rand_str(str_size, alphanum)
}

func RandInt(str_size int) int {

	alphanum := rand_seek(Number)
	return rand_int(str_size, alphanum)
}

func rand_str(str_size int, alphanum string) string {
	var bytes = make([]byte, str_size)
	// put here it's very slow, go put init !!!!
	//rand.Seed(time.Now().UTC().UnixNano()) //very important
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func rand_int(str_size int, alphanum string) int {
	var bytes = make([]byte, str_size)
	noZero := rand_seek(NumberNoZero)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	for bytes[0] == 0 {
		bytes[0] = noZero[rand.Int()%(len(noZero))]
	}

	r, err := strconv.Atoi(string(bytes))
	if err != nil {
		panic(err)
		return 0
	}
	return r

}
