package environment

import (
	"flag"
	"log"
)

type Flag struct {
	SpecialSemester bool
}

const (
	FlagRegister   = "register"
	FlagShow       = "show"
	FlagIntrospect = "introspect"
)

var validCmd = []string{
	FlagRegister,
	FlagShow,
	FlagIntrospect,
}

func CheckArgs(osArgs []string) {
	cmdExists := false
	for i := range validCmd {
		if validCmd[i] == osArgs[1] {
			cmdExists = true
		}
	}
	if !cmdExists {
		log.Fatalf("Unknown argument: '%v'\n", osArgs[1])
	}
}

func (env *Environment) RetrieveCommandFlag(args []string) {
	if args[1] == FlagRegister {
		flagSet := flag.NewFlagSet(ProjectName, flag.PanicOnError)
		flagSet.BoolVar(&env.Flag.SpecialSemester, "special-semester", false, "Register the semester 0 as a valid one.")
		_ = flagSet.Parse(args[2:])
	} else if args[1] == FlagIntrospect {
		env.Flag.SpecialSemester = true
	}
}
