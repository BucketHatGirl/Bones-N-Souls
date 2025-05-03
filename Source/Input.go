package main
import (
	"syscall/js"
)

var EVENTS = make(map[string]bool)

func RegisterKeyboard() {
	D := GetGlobal("window")
	D.Call("addEventListener", "keydown", js.FuncOf(keyDown))
	D.Call("addEventListener", "keyup", js.FuncOf(keyUp))
}

func keyDown(this js.Value, D []js.Value) interface{} {
	EVENTS[D[0].Get("key").String()] = true
	return nil
}

func keyUp(this js.Value, U []js.Value) interface{} {
	EVENTS[U[0].Get("key").String()] = false
	return nil
}
