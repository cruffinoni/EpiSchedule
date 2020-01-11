package environment

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Flag struct {
	SpecialSemester bool
}

type FlagType []string
type FlagArg struct {
	Hold         interface{}
	DefaultValue interface{}
	Name         string
	Description  string
}

const (
	FlagRegister   = "register"
	FlagShow       = "show"
	FlagIntrospect = "introspect"
)

var FlagArgEmpty FlagArg // []FlagArgEmpty

var (
	validCmd = FlagType{
		FlagRegister,
		FlagShow,
		FlagIntrospect,
	}

	cmdArg = map[string]FlagArg {
		FlagRegister: FlagArgEmpty,
		FlagShow: FlagArgEmpty,
		FlagIntrospect: FlagArgEmpty,
	}
)

func (i *FlagType) String() string {
	return "my string representation"
}

func (i *FlagType) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func SetArgToCmd(cmd string, arg FlagArg) {
	if cmdArg[cmd] != FlagArgEmpty {
		return
	}
	cmdArg[cmd] = arg
}

func (env *Environment) RetrieveCommandFlag(args []string) {
	flagSet := flag.NewFlagSet(ProjectName, flag.PanicOnError)
	flagSet.Var(&validCmd, "command", "The main command")
	if len(args) < 2 {
		flagSet.Usage()
		os.Exit(1)
	}
	_ = flagSet.Parse(args[2:])
	for cmd, arg := range cmdArg {
		if arg != FlagArgEmpty {
			argCmdSet := flag.NewFlagSet(cmd, flag.PanicOnError)
			fmt.Printf("New arg for cmd: %v w/ %v\n", cmd, arg.Hold)
			switch arg.Hold.(type) {
			case nil:
				log.Fatalf("command %v has an nil type\n", arg.Name)
			case *int:
				arg.Hold = argCmdSet.Int(arg.Name, arg.DefaultValue.(int), arg.Description)
			case *bool:
				arg.Hold = argCmdSet.Bool(arg.Name, arg.DefaultValue.(bool), arg.Description)
			case *string:
				arg.Hold = argCmdSet.String(arg.Name, arg.DefaultValue.(string), arg.Description)
			default:
				log.Fatalf("command %v has an unknown type\n", arg.Name)
			}
			//if len(args) < 3 {
			//	argCmdSet.Usage()
			//	os.Exit(1)
			//}
			_ = flagSet.Parse(args[2:])
		}
	}
}

func InitCommandArg(env *Environment) {
	SetArgToCmd(FlagRegister, FlagArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: false,
		Name:         "special-semester",
		Description:  "Register the semester 0 as a valid one.",
	})
	SetArgToCmd(FlagIntrospect, FlagArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: true,
		Name:         "special-semester",
		Description:  "Register the semester 0 as a valid one. It will give more type.",
	})
}