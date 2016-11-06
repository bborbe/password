package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestServerSuccess(t *testing.T) {
	srv, err := createServer()
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(srv, NotNilValue()); err != nil {
		t.Fatal(err)
	}
}
