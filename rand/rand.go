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

const (
	number        = "0123456789"
	upperCharater = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCharater = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano()) //very important
}

func randSeek(randType int) string {

	switch randType {
	case Number:
		return number
	case NumberNoZero:
		return "123456789"
	case Hex:
		return number + "ABCDEF"
	case Upper:
		return upperCharater
	case Lower:
		return lowerCharater
	case MixAll:
		return number + upperCharater + lowerCharater
	case MixUpper:
		return number + upperCharater
	case MixLower:
		return number + lowerCharater
	default:
		return upperCharater + lowerCharater
	}
}

func RandString(str_size int, rand_type ...int) string {

	mix := Mix
	if len(rand_type) > 0 {
		mix = rand_type[0]
	}
	alphanum := randSeek(mix)
	return randStr(str_size, alphanum)
}

func RandInt(str_size int) int {

	alphanum := randSeek(Number)
	return randInt(str_size, alphanum)
}

func randStr(str_size int, alphanum string) string {
	var bytes = make([]byte, str_size)
	// put here it's very slow, go put init !!!!
	//rand.Seed(time.Now().UTC().UnixNano()) //very important
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func randInt(str_size int, alphanum string) int {
	var bytes = make([]byte, str_size)
	noZero := randSeek(NumberNoZero)
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
