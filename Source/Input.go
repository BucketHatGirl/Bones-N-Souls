package main
import (
	"syscall/js"
)
type Input struct {
	EVENTS map[string]bool
}

func (I Input) KeyDown(this js.Value, KEY []js.Value) interface{} {
	if KEY[0].Length() > 0 {
		I.EVENTS[KEY[0].Get("key").String()] = true
	}
	return nil
}

func (I Input) KeyUp(this js.Value, KEY []js.Value) interface{} {
	if KEY[0].Length() > 0 {
		I.EVENTS[KEY[0].Get("key").String()] = false
	}
	return nil
}

func (I Input) RegisterKeyboard() {
	var DOCUMENT = new(Webpage).GetGlobal("window")
	DOCUMENT.Call("addEventListener", "keydown", js.FuncOf(I.KeyDown))
	DOCUMENT.Call("addEventListener", "keyup", js.FuncOf(I.KeyUp))
}

