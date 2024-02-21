package rand

import (
	"math/rand"
	"strconv"
)

func Rander() string {
	randomNumber := rand.Intn(10)
	Randnum := strconv.Itoa(randomNumber)

	return Randnum
}
