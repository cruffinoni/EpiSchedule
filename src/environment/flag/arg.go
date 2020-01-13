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

type HandlerType func(environment.Environment, []blueprint.Course)
type PreHandlerType func(*environment.Environment)

type ProgArg struct {
	Hold         interface{}
	DefaultValue interface{} // Set a default value if the arg is not optional
	Name         string
	Description  string
}

type ProgCmd struct {
	Args       []ProgArg
	preHandler PreHandlerType
	handler    HandlerType
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
