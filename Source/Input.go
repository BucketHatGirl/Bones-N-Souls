package main
import (
	"syscall/js"
)

type Input struct {
	EVENT []string
} 

func (I Input) Poll() []string {
	I.EVENT = make([]string, 32)
	var KEYS = js.Global().Call("FetchKeys")
	if !KEYS.IsUndefined() {
		for X := 0; X < KEYS.Length(); X++ {
    	I.EVENT[X] = KEYS.Index(X).String()
  	}
	}
	return I.EVENT
}
