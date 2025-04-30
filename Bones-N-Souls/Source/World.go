package main
import (
	"syscall/js"
)

var CANNON js.Value = GetGlobal("CANNON")
var WORLD js.Value = CANNON.Get("World").New()
var INDEX int = 1

func InitWorld() {
	WORLD.Get("gravity").Call("set", 0, -9.81, 0)
}

type Object struct {
	ID int
	VELOCITY []float64
	POSITION []float64
	ROTATION []float64
	DATA []interface{}
}

func NewObject() (*Object, interface{}) {
	if INDEX < 32000 {
		X := new(Object)
		X.ID = INDEX
		INDEX++
		return X, nil
	}
	return nil, nil
}

