package main
import (
	"syscall/js"
)

var CANNON js.Value = GetGlobal("CANNON")
var WORLD js.Value = CANNON.Get("World").New()
var INDEX int = 1

type Object struct {
	OBJECT js.Value
	BODY js.Value
	POSITION [3]int
	ROTATION [3]int
}

func SpawnObject(OBJ *Object) {
	OBJ.BODY.Call("addShape", OBJ.OBJECT)
	WORLD.Call("addBody", OBJ.BODY)
}

func InitWorld() {
	WORLD.Get("gravity").Call("set", 0, -9.81, 0)
	GROUND := new(Object)
	GROUND.OBJECT = CANNON.Get("Plane").New()
	GROUND.OBJECT.Set("mass", 0)
	GROUND.POSITION = [3]int{0, 0, 0}
	GROUND.ROTATION = [3]int{90, 0, 0}
	GROUND.BODY = CANNON.Get("Body").New()
	SpawnObject(GROUND)
}
