package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"runtime"

	"github.com/bborbe/log"
	password_generator "github.com/bborbe/password/generator"
)

const (
	DEFAULT_PASSWORD_LENGTH   = 16
	PARAMETER_PASSWORD_LENGTH = "length"
	PARAMETER_LOGLEVEL        = "loglevel"
)

var (
	logger            = log.DefaultLogger
	logLevelPtr       = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	passwordLengthPtr = flag.Int(PARAMETER_PASSWORD_LENGTH, DEFAULT_PASSWORD_LENGTH, "string")
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	writer := os.Stdout
	passwordGenerator := password_generator.New()
	err := do(writer, passwordGenerator, *passwordLengthPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, passwordGenerator password_generator.PasswordGenerator, passwordLength int) error {
	logger.Debug("start")
	if passwordLength < 0 {
		fmt.Fprintf(writer, "illegal password length\n")
		return nil
	}
	fmt.Fprintf(writer, "%s\n", passwordGenerator.GeneratePassword(passwordLength))
	logger.Debug("done")
	return nil
}
