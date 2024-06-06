package recovery

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func Recover(logger ...*logrus.Logger) {
	if err := recover(); nil != err {
		PrintStackTrace(err, logger...)
	}
}

// StackTrace return error stack
func StackTrace(e interface{}) error {
	stack := make([]byte, 1024*8)
	stack = stack[:runtime.Stack(stack, false)]
	stackTraces := fmt.Sprintf("panic: %v\n%s\n", e, stack)

	return errors.New(stackTraces)
}

// PrintStackTrace print stack trace
func PrintStackTrace(e interface{}, logger ...*logrus.Logger) {
	if len(logger) > 0 {
		logger[0].Println(StackTrace(e))
	}
	fmt.Println(StackTrace(e))
}
