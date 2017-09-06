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

//RandString generate rand string
func RandString(strSize int, randType ...int) string {

	mix := Mix
	if len(randType) > 0 {
		mix = randType[0]
	}
	alphanum := randSeek(mix)
	return randStr(strSize, alphanum)
}

//RandInt generate rand int
func RandInt(strSize int) int {

	alphanum := randSeek(Number)
	return randInt(strSize, alphanum)
}

//randStr package internal func
func randStr(strSize int, alphanum string) string {
	var bytes = make([]byte, strSize)
	// put here it's very slow, go put init !!!!
	//rand.Seed(time.Now().UTC().UnixNano()) //very important
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

//randInt package internal func
func randInt(strSize int, alphanum string) int {
	var bytes = make([]byte, strSize)
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
	}
	return r

}
