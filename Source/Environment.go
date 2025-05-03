package main
import (
	"syscall/js"
)

type Environment struct {
	DATA map[string]js.Value
}

func (E Environment) NewEnvironment() {
	var WEBPAGE = new(Webpage)
	E.DATA = make(map[string]js.Value, 2)
	E.DATA["CANNON"] = WEBPAGE.GetGlobal("CANNON")
	E.DATA["WORLD"] = E.DATA["CANNON"].Get("World").New()
}


