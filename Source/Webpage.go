package main
import (
	"syscall/js"
)
type Webpage struct {}

func (W Webpage) GetGlobal(NAME string) js.Value {
	return js.Global().Get(NAME)
}

func (W Webpage) SetGlobal(NAME string, VALUE interface{}) {
	js.Global().Set(NAME, VALUE)
}

func (W Webpage) Evaluate(CODE string) js.Value {
	return W.GetGlobal("eval").Invoke(CODE)
}
