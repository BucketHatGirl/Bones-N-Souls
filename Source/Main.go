package main

func main() {
	var INPUT = new(Input)
	var WEBPAGE = new(Webpage)
	var ENVIRONMENT = new(Environment)
	var RENDERER = new(Renderer)
	var CONSOLE = WEBPAGE.GetGlobal("console")
	CONSOLE.Call("log", `"Bones 'N Souls" has been loaded.`)
	var FOV int = 75
	INPUT.RegisterKeyboard()
	ENVIRONMENT.NewEnvironment()
	RENDERER.NewRenderer()
	RENDERER.NewCamera(0, FOV)
	go func() {
		for {
				
		}
	}()
	select {}
}
