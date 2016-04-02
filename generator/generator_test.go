package generator

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsPasswordGenerator(t *testing.T) {
	b := New()
	var i *PasswordGenerator
	err := AssertThat(b, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerate(t *testing.T) {
	b := New()
	for i := 0; i < 32; i++ {
		if err := AssertThat(len(b.GeneratePassword(i)), Is(i)); err != nil {
			t.Fatal(err)
		}
	}
}
func TestGenerateMaxLength(t *testing.T) {
	b := New()
	if err := AssertThat(len(b.GeneratePassword(MAX_LENGTH*2)), Is(MAX_LENGTH)); err != nil {
		t.Fatal(err)
	}
}
