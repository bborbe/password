package handler

import (
	"net/http"
	password_generator "github.com/bborbe/password/generator"
	"github.com/bborbe/log"
	"fmt"
	"strconv"
)

var logger = log.DefaultLogger

const (
	DEFAULT_LENGTH = 16
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
	logger.Debugf("generate password with length %d", length)
	password := s.passwordGenerator.GeneratePassword(length)
	fmt.Fprintf(responseWriter, password)
}

func parseLength(lengthString string) int {
	i, err := strconv.Atoi(lengthString)
	if err != nil {
		return DEFAULT_LENGTH
	}
	return i
}
