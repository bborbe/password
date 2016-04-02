package generator

import (
	"time"
	"math/rand"
)

var (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NUMBERS = []rune("0123456789")
	SPECIAL = []rune("!@#$%&;,.")
	ALL = append(append(LETTERS, NUMBERS...), SPECIAL...)
)

type passwordGenerator struct {

}

type PasswordGenerator interface {
	GeneratePassword(length int) (string)
}

func New() *passwordGenerator {
	return new(passwordGenerator)
}

func (p *passwordGenerator) GeneratePassword(length int) (string) {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	pos := 0
	for i := range b {
		pos++
		b[i] = random(pos)
	}
	return string(b)
}

func random(pos int) rune {
	if pos == 1 {
		return LETTERS[rand.Intn(len(LETTERS))]
	}
	return ALL[rand.Intn(len(ALL))]
}
