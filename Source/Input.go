package main
import (
	"syscall/js"
)
type Input struct {
	EVENTS map[string]bool
}

func (I Input) KeyDown(this js.Value, KEY []js.Value) interface{} {
	I.EVENTS[KEY[0].Get("key").String()] = true
	return nil
}

func (I Input) KeyUp(this js.Value, KEY []js.Value) interface{} {
	I.EVENTS[KEY[0].Get("key").String()] = false
	return nil
}

func (I Input) RegisterKeyboard() {
	W := new(Webpage)
	D := W.GetGlobal("window")
	D.Call("addEventListener", "keydown", js.FuncOf(I.KeyDown))
	D.Call("addEventListener", "keyup", js.FuncOf(I.KeyUp))
}

