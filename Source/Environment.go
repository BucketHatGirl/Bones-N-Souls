package main
import (
	"syscall/js"
)

type Environment struct {
	DATA map[string]js.Value
}

func (E Environment) NewEnvironment() {
	E.DATA["CANNON"] = new(Webpage).GetGlobal("CANNON")
	E.DATA["WORLD"] = E.DATA["CANNON"].Get("World").New()
}


