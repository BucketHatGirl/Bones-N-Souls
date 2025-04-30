package main
import (
	"syscall/js"
)

var WINDOW js.Value = GetGlobal("window")
var DOCUMENT js.Value = GetGlobal("document")
var THREE js.Value = GetGlobal("THREE")
var SCENE js.Value = THREE.Get("Scene").New()
var RENDERER js.Value = THREE.Get("WebGLRenderer").New()
var HEIGHT int = WINDOW.Get("innerHeight").Int()
var WIDTH int = WINDOW.Get("innerWidth").Int()
var CAMERA js.Value = THREE.Get("PerspectiveCamera").New(75, WIDTH / HEIGHT, 0.1, 1000)

func InitRenderer() {
  RENDERER.Call("setSize", WIDTH, HEIGHT)
  DOCUMENT.Get("body").Call("appendChild", RENDERER.Get("domElement"))
}

