package generator

import "bytes"

type passwordGenerator struct {

}

type PasswordGenerator interface {
	GeneratePassword(length int) (string)
}

func New() *passwordGenerator {
	return new(passwordGenerator)
}

func (p *passwordGenerator) GeneratePassword(length int) (string) {
	buffer := bytes.NewBufferString("")
	for i := 0; i < length; i++ {
		buffer.WriteString("a")
	}
	return buffer.String()
}