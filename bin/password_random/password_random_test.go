package main

import (
	"testing"

	"bytes"

	password_generator "github.com/bborbe/password/generator"

	. "github.com/bborbe/assert"
)

func TestResumeFail(t *testing.T) {
	writer := bytes.NewBufferString("")
	passwordGenerator := password_generator.New()

	err := do(writer, passwordGenerator, -1)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("illegal password length\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestResumeSuccess(t *testing.T) {
	writer := bytes.NewBufferString("")
	passwordGenerator := password_generator.New()
	length := 14
	err := do(writer, passwordGenerator, length)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(writer.String()), Is(length+1))
	if err != nil {
		t.Fatal(err)
	}
}
