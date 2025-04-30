package main
import (
	"syscall/js"
	"time"
)
var PRESSED js.Value 
var CONSOLE js.Value = GetGlobal("console")
func Runtime() {
	var TIME int
	go func() {
		for {
			TIME++
			time.Sleep(1000)
		}
	}()
	RegisterKeyboard()
	go func() {
		for {
			PRESSED = PollKeyboard()
		}
	}()
	InitRenderer()
	InitWorld()
}

func main() {
	GetGlobal("console").Call("log","Bones 'N Souls has been loaded.")
	C := make(chan int)
	Runtime()
	<-C
}
