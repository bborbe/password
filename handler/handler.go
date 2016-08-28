package handler

import (
	"fmt"
	"net/http"
	"strconv"

	password_generator "github.com/bborbe/password/generator"
	"github.com/golang/glog"
)

const (
	DEFAULT_LENGTH   = 16
	PARAMETER_LENGTH = "length"
)

type statusHandler struct {
	passwordGenerator password_generator.PasswordGenerator
}

func New(passwordGenerator password_generator.PasswordGenerator) http.Handler {
	s := new(statusHandler)
	s.passwordGenerator = passwordGenerator
	return s
}

func (s *statusHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	length := parseLength(request.FormValue(PARAMETER_LENGTH))
	glog.V(2).Infof("generate password with length %d", length)
	password := s.passwordGenerator.GeneratePassword(length)
	fmt.Fprint(responseWriter, password)
}

func parseLength(lengthString string) int {
	i, err := strconv.Atoi(lengthString)
	if err != nil {
		return DEFAULT_LENGTH
	}
	return i
}
