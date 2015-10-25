package strings

import (
	"bytes"
	"math/rand"
	"time"
)

type Class interface {
	Make() string
}

type RandomPicker struct {
	length uint
	runes  []rune
	random *rand.Rand
}

func (r RandomPicker) Make() string {
	var buffer bytes.Buffer
	for i := uint(0); i < r.length; i++ {
		buffer.WriteRune(r.runes[r.random.Intn(len(r.runes))])
	}
	return buffer.String()
}

func NewRandomPicker(length uint, runes []rune) RandomPicker {
	return RandomPicker{
		length: length,
		runes:  runes,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
