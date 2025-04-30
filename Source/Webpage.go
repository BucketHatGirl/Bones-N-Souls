package main
import (
	"syscall/js"
)

func GetGlobal(NAME string) js.Value {
	return js.Global().Get(NAME)
}

func SetGlobal(NAME string, VALUE interface{}) {
	js.Global().Set(NAME, VALUE)
}

