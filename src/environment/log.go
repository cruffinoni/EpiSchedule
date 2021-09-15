package environment

import (
	"fmt"
	"log"
	"time"
)

type VerboseLevel uint

const (
	VerboseSimple  = iota + 1
	VerboseMedium  = VerboseSimple + 1
	VerboseDebug   = VerboseMedium + 1
	VerboseDefault = VerboseDebug
)

var verboseLevelName = map[VerboseLevel]string{
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
		msg = time.Now().Format(time.RFC3339) + " " + msg
		fmt.Printf(msg+ColorReset, format...)
	}
}

func (env Environment) Log(level VerboseLevel, msg ...interface{}) {
	if env.verbose >= level {
		fmt.Print(time.Now().Format(time.RFC3339) + " ")
		fmt.Print(msg...)
		fmt.Print(ColorReset)
	}
}

func (env Environment) Error(msg ...interface{}) {
	log.Print(msg...)
}

func (env Environment) Errorf(msg string, fmt ...interface{}) {
	log.Printf(msg, fmt...)
}

func (env Environment) Fatal() {
	log.Panic()
}
