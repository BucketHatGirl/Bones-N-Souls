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

func Evaluate(CODE string) js.Value {
	return GetGlobal("eval").Invoke(CODE)
}
