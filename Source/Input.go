package main
import (
	"syscall/js"
)

var EVENTS = make(chan js.Value)

func RegisterKeyboard() {
	D := GetGlobal("document")
	D.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, X []js.Value) interface{} {
		EVENTS <- X[0]
		return nil
	}))
}

func PollKeyboard() string {
	EVENT := <-EVENTS
	return EVENT.Get("key").String()
}
