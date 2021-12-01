package wordz

import (
	"math/rand"
	"time"
)

func genRandLetter() string {
	return string(rand.Intn(123 - 97) + 97)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func GenWord() string {
	length := random(3, 10)
	word := ""
	for i:= 0; i < length; i++ {
		word += genRandLetter()
	}
	return word
}

func init() {
	var SEED = time.Now().Unix()
	rand.Seed(SEED)
}