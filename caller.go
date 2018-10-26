package callertest

import (
	"runtime"
)

func Call() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func IndirectCall() string {
	return Call()
}
