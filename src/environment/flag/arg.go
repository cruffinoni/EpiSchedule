package flag

import (
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
)

const (
	ArgRegister   = "register"
	ArgShow       = "show"
	ArgIntrospect = "introspect"

	stringDefaultValue = ""
	intDefaultValue    = 0
	boolDefaultValue   = false
)

type HandlerCmd func(environment.Environment, []blueprint.Course)

type ProgArg struct {
	Hold         interface{}
	DefaultValue interface{} // Set a default value if the arg is not optional
	Name         string
	Description  string
}

type ProgCmd struct {
	Args       ProgArg // []ProgArg
	preHandler func()
	handler    HandlerCmd
}

func (ref ProgArg) IsEqual(toCompare ProgArg) bool {
	if ref.DefaultValue == toCompare.DefaultValue &&
		ref.Description == toCompare.Description &&
		ref.Name == toCompare.Name &&
		ref.Hold == toCompare.Hold {
		return true
	}
	return false
}
