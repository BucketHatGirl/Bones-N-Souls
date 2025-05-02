package main
import (
	"syscall/js"
)

var CONSOLE js.Value = GetGlobal("console")
var INPUT string
func main() {
	CONSOLE.Call("log", "Program is loaded")
	var SPEEDCAP [2]float64 = [2]float64{4.4, -4.4}
	var FORWARD_VELOCITY float64 = 0.0
	var SIDWAYS_VELOCITY float64 = 0.0
	RegisterKeyboard()
	go func() {
		for {
			INPUT = PollKeyboard()
			if INPUT != "" {
				switch INPUT {
					case "w":
						if FORWARD_VELOCITY < SPEEDCAP[0] {
							FORWARD_VELOCITY = FORWARD_VELOCITY + 1.1
						}
					case "a":
						if SIDWAYS_VELOCITY > SPEEDCAP[1] {
							SIDWAYS_VELOCITY = SIDWAYS_VELOCITY - 1.1
						}
					case "s":
						if FORWARD_VELOCITY > SPEEDCAP[1] {
							FORWARD_VELOCITY = FORWARD_VELOCITY - 1.1
						}
					case "d":
						if SIDWAYS_VELOCITY < SPEEDCAP[0] {
							SIDWAYS_VELOCITY = SIDWAYS_VELOCITY + 1.1
						}
				}
			}
		}
	}()
	select {}
}
