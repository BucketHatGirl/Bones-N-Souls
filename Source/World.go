package main
import (
	"syscall/js"
)

var WINDOW js.Value = GetGlobal("window")
var DOCUMENT js.Value = GetGlobal("document")
var THREE js.Value = GetGlobal("THREE")
var SCENE js.Value = THREE.Get("Scene").New()
var HEIGHT int = WINDOW.Get("innerHeight").Int()
var WIDTH int = WINDOW.Get("innerWidth").Int()
var CAMERA js.Value = THREE.Get("PerspectiveCamera").New(75, WIDTH / HEIGHT, 0.1, 1000)
var RENDERER js.Value = THREE.Get("WebGLRenderer").New()
var CANNON js.Value = GetGlobal("CANNON")
var WORLD js.Value = CANNON.Get("World").New()

func Init() {
	WORLD.Get("gravity").Call("set", 0, -9.81, 0)

}

type Object struct {
	ID int
	OBJECT js.Value
	BODY js.Value
	POSITION [3]int
	ROTATION [3]int
}

func NewObject(ID int) *Object {
	OBJECT := new(Object)
	OBJECT.ID = ID
	return OBJECT
}

func RenderObject() {
	
}
