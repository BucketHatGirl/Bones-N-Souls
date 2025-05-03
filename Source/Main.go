package main

func main() {
	var INPUT = new(Input)
	var WEBPAGE = new(Webpage)
	var ENVIRONMENT = new(Environment)
	var RENDERER = new(Renderer)
	var CONSOLE = WEBPAGE.GetGlobal("console")
	CONSOLE.Call("log", `"Bones 'N Souls" has been loaded.`)
	INPUT.RegisterKeyboard()
	ENVIRONMENT.NewEnvironment()
	RENDERER.NewRenderer()
	go func() {
		for {
			if INPUT.EVENTS["w"] {
				CONSOLE.Call("log", 1)
			}	
		}
	}()
	select {}
}
