package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"runtime"

	password_generator "github.com/bborbe/password/generator"
	"github.com/golang/glog"
)

const (
	DEFAULT_PASSWORD_LENGTH   = 16
	PARAMETER_PASSWORD_LENGTH = "length"
)

var (
	passwordLengthPtr = flag.Int(PARAMETER_PASSWORD_LENGTH, DEFAULT_PASSWORD_LENGTH, "string")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	writer := os.Stdout
	passwordGenerator := password_generator.New()
	err := do(writer, passwordGenerator, *passwordLengthPtr)
	if err != nil {
		glog.Exit(err)
	}
}

func do(writer io.Writer, passwordGenerator password_generator.PasswordGenerator, passwordLength int) error {
	glog.V(2).Info("start")
	if passwordLength < 0 {
		fmt.Fprintf(writer, "illegal password length\n")
		return nil
	}
	fmt.Fprintf(writer, "%s\n", passwordGenerator.GeneratePassword(passwordLength))
	glog.V(2).Info("done")
	return nil
}
