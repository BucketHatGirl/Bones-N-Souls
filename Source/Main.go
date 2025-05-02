package main
import (
	"syscall/js"
	"time"
	"strings"
)
var PRESSED interface{} 
var CONSOLE js.Value = GetGlobal("console")
func Runtime() {
	var VELOCITY [3]float64
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
	PRESSED = strings.ToLower(PRESSED.String())	
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
