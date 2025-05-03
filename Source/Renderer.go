package main
import (
	"syscall/js"
)

var WINDOW js.Value = new(Webpage).GetGlobal("window")
var HEIGHT int = WINDOW.Get("innerHeight").Int()
var WIDTH int = WINDOW.Get("innerWidth").Int()

type Renderer struct {
	DATA map[string]js.Value
	CAMERA []js.Value
}

func (R Renderer) NewRenderer() {
	R.DATA = make(map[string]js.Value)
	R.DATA["THREE"] = new(Webpage).GetGlobal("THREE")
	R.DATA["SCENE"] = R.DATA["THREE"].Get("Scene").New()
	R.DATA["RENDERER"] = R.DATA["THREE"].Get("WebGLRenderer").New()
}

func (R Renderer) NewCamera(INDEX int, FOV int) {
	R.CAMERA[INDEX] = R.DATA["THREE"].Get("PerspectiveCamera").New(FOV, WIDTH / HEIGHT, 0.1, 1000)
}

func (R Renderer) Init() {

}

