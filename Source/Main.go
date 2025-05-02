package main
import (
	"syscall/js"
	"time"
)
var PRESSED string 
var CONSOLE js.Value = GetGlobal("console")
func Runtime() {
	var SPEEDCAP float64 = 3.2
	var FORWARD_VELOCITY float64 = 0.0
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
				if FORWARD_VELOCITY < SPEEDCAP {
					FORWARD_VELOCITY = FORWARD_VELOCITY + 0.1
				}	
			} else {
				if FORWARD_VELOCITY > 0 {
					FORWARD_VELOCITY = FORWARD_VELOCITY - 0.1
				}
			}
		}
	}()
	go func() {
		for {
			if FORWARD_VELOCITY > 0 {
				CONSOLE.Call("log", js.ValueOf(FORWARD_VELOCITY))
			}	
		}
	}()

}

func main() {
	GetGlobal("console").Call("log","Bones 'N Souls has been loaded.")
	C := make(chan int)
	Runtime()
	<-C
}
