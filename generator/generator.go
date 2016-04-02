package generator

import (
	"math/rand"
	"time"
)

var (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NUMBERS = []rune("0123456789")
	SPECIAL = []rune("!@#$%&;,.")
	ALL     = append(append(LETTERS, NUMBERS...), SPECIAL...)
)

const MAX_LENGTH = 64

type passwordGenerator struct {
}

type PasswordGenerator interface {
	GeneratePassword(length int) string
}

func New() *passwordGenerator {
	return new(passwordGenerator)
}

func (p *passwordGenerator) GeneratePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, min(length, MAX_LENGTH))
	pos := 0
	for i := range b {
		pos++
		b[i] = random(pos)
	}
	return string(b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func random(pos int) rune {
	if pos == 1 {
		return LETTERS[rand.Intn(len(LETTERS))]
	}
	return ALL[rand.Intn(len(ALL))]
}
