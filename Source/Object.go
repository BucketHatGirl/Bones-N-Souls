package main
import (
	"syscall/js"
)

type Object struct {
	ID int
	OBJECT map[string]js.Value
	ROTATiON [3]float64
	POSITION [3]float64
	VELOCITY [3]float64
}
 
