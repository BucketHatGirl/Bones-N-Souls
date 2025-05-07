package main
import (
	"embed"
	"time"
)
//go:embed Javascript/*
var JAVASCRIPT embed.FS
//go:embed Assets/*
var ASSETS embed.FS

func init() {
	var WEBPAGE = new(Webpage)
	var FILES []string = []string{"Input.js", "Cannon.js", "Three.js"}
	for _, X := range FILES {
		var CONTENT, _ = JAVASCRIPT.ReadFile(X)
		var SCRIPT = WEBPAGE.GetGlobal("document").Call("createElement", "script")
		SCRIPT.Set("type", "text/javascript")
		SCRIPT.Set("innerHTML", string(CONTENT))
		WEBPAGE.GetGlobal("document").Get("head").Call("appendChild", SCRIPT)
	}
}

func main() {
	var WEBPAGE = new(Webpage)
	var INPUT = new(Input)
	var CONSOLE = WEBPAGE.GetGlobal("console")
	var ENVIRONMENT = new(Environment)
 	var RENDERER = new(Renderer)	
	ENVIRONMENT.NewEnvironment()
	RENDERER.NewRenderer()
	go func() {
		for {
			var EVENTS = INPUT.Poll()
			for _, E := range EVENTS {
				if E != "" {
					switch E {
						case "w":
							
						case "a":
			
						case "s":

						case "d":

						case "Shift":

						case "Control":

						case "Alt":

					}
				}
			}
			time.Sleep(time.Second / 7)
		}
	}()
	CONSOLE.Call("log", `"Bones 'N Souls" Runtime has started.`)
	select {}
}
