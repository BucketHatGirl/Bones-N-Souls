package main
import (
	"syscall/js"
)

var LIMIT int = 3200
var OBJECT_STORAGE = make([]interface{}, LIMIT)

type Object struct {
	ID int
	COLLIDER js.Value
	STYLE js.Value
	DATA map[string]bool
}

func (O Object) NewObject() *Object {
	return new(Object)
}

 
