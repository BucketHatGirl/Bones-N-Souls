package main

func main() {
	var INPUT = new(Input)
	var WEBPAGE = new(Webpage) 
	var CONSOLE = WEBPAGE.GetGlobal("console")
	CONSOLE.Call("log", `"Bones 'N Souls" has been loaded.`)
	INPUT.RegisterKeyboard()
	go func() {
		for {
			EVENTS := INPUT.EVENTS
			if EVENTS["w"] {CONSOLE.Call("log", "w")}
		}
	}()
	select {}
}
