package environment

import (
	"fmt"
	"log"
)

type VerboseLevel uint

const (
	VerboseNone    VerboseLevel = iota
	VerboseSimple               = VerboseNone + 1
	VerboseMedium               = VerboseSimple + 1
	VerboseDebug                = VerboseMedium + 1
	VerboseDefault              = VerboseDebug
)

var verboseLevelName = map[VerboseLevel]string{
	VerboseNone:   "None",
	VerboseSimple: "Simple",
	VerboseMedium: "Medium",
	VerboseDebug:  "Debug",
}

func (env Environment) SetVerboseLevel(level VerboseLevel) {
	if verboseLevelName[level] == "" {
		return
	}
	env.verbose = level
}

func (env Environment) GetVerboseLevel() VerboseLevel {
	return env.verbose
}

func (env Environment) Logf(level VerboseLevel, msg string, format ...interface{}) {
	if env.verbose >= level {
		fmt.Printf(msg, format...)
	}
}

func (env Environment) Log(level VerboseLevel, msg ...interface{}) {
	if env.verbose >= level {
		fmt.Print(msg...)
	}
}

func (env Environment) Error(msg ...interface{}) {
	log.Print(msg...)
}

func (env Environment) Errorf(msg string, fmt ...interface{}) {
	log.Printf(msg, fmt...)
}
