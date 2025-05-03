package main
import (
	"syscall/js"
)

var CONSOLE js.Value = GetGlobal("console")

func main() {
	CONSOLE.Call("log", "Program is loaded")
	RegisterKeyboard()
	go func() {
		for {
				
		}
	}()
	Init()
	select {}
}
