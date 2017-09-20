package crypto

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

//Sha1 provide simple []byte/string sha1
func Sha1(input []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(input))
}

//Sha1File provide simple file sha1 sum
func Sha1File(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil

}
