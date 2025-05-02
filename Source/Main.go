package main
import (
	"syscall/js"
	"time"
	"strings"
)
var PRESSED string 
var CONSOLE js.Value = GetGlobal("console")
func Runtime() {
	var VELOCITY []float64
	var TIME int = 0
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
	go func() {
		for {
			if PRESSED == "w" {
				
			}
		}
	}()
	go func() {
		for {
		
		}
	}()

}

func main() {
	GetGlobal("console").Call("log","Bones 'N Souls has been loaded.")
	C := make(chan int)
	Runtime()
	<-C
}
