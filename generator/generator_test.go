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
