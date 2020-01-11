package environment

import (
	"flag"
	"log"
	"os"
	"strings"
)

// Env flag type
type Flag struct {
	SpecialSemester bool
}

type FlagType []string
type FlagArg struct {
	Hold         interface{}
	DefaultValue interface{} // Set a default value if the arg is not optional
	Name         string
	Description  string
}

const (
	FlagRegister   = "register"
	FlagShow       = "show"
	FlagIntrospect = "introspect"

	stringDefaultValue = ""
	intDefaultValue    = 0
	boolDefaultValue   = false
)

var FlagArgEmpty FlagArg // []FlagArgEmpty

var (
	validCmd = FlagType{
		FlagRegister,
		FlagShow,
		FlagIntrospect,
	}

	cmdArg = map[string]FlagArg{
		FlagRegister:   FlagArgEmpty,
		FlagShow:       FlagArgEmpty,
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

func SetArgToCmd(env Environment, cmd string, arg FlagArg) {
	if cmdArg[cmd] != FlagArgEmpty {
		env.Errorf("Cmd %v has already args set\n", cmd)
		return
	}
	cmdArg[cmd] = arg
}

func printProgramUsage(env Environment, optional ...string) {
	if len(optional) < 3 {
		env.Errorf("Usage of %v:\n\t./%v [command] [arguments]\n\t\t"+
			"command: One of these main commands: %v\n\t\t"+
			"arguments: The argument(s) associated to the command\n",
			ProjectName, strings.ToLower(ProjectName), validCmd)
	} else {
		env.Errorf("Usage of %v:\n\t./%v %v [arguments]\n\t\t"+
			"%v: %v\n",
			ProjectName, strings.ToLower(ProjectName), optional[0],
			optional[1], optional[2])
	}
	os.Exit(1)
}

func (env *Environment) RetrieveCommandFlag(args []string) {
	if len(args) < 2 {
		printProgramUsage(*env)
	}
	for cmd, arg := range cmdArg {
		if arg != FlagArgEmpty && args[1] == cmd {
			argCmdSet := flag.NewFlagSet(cmd, flag.ExitOnError)
			switch arg.Hold.(type) {
			case nil:
				log.Fatalf("command %v has an nil type\n", arg.Name)
			case *int:
				defaultValue := intDefaultValue
				if arg.DefaultValue != nil {
					defaultValue = arg.DefaultValue.(int)
				}
				argCmdSet.IntVar(arg.Hold.(*int), arg.Name, defaultValue, arg.Description)
			case *bool:
				defaultValue := boolDefaultValue
				if arg.DefaultValue != nil {
					defaultValue = arg.DefaultValue.(bool)
				}
				argCmdSet.BoolVar(arg.Hold.(*bool), arg.Name, defaultValue, arg.Description)
			case *string:
				defaultValue := stringDefaultValue
				if arg.DefaultValue != nil {
					defaultValue = arg.DefaultValue.(string)
				}
				argCmdSet.StringVar(arg.Hold.(*string), arg.Name, defaultValue, arg.Description)
			default:
				log.Fatalf("command %v has an unknown type\n", arg.Name)
			}
			_ = argCmdSet.Parse(args[2:])
			switch arg.Hold.(type) {
			case *int:
				if arg.DefaultValue == nil && *arg.Hold.(*int) == intDefaultValue {
					printProgramUsage(*env, cmd, arg.Name, arg.Description)
				}
			case *bool:
				if arg.DefaultValue == nil && *arg.Hold.(*bool) == boolDefaultValue {
					printProgramUsage(*env, cmd, arg.Name, arg.Description)
				}
			case *string:
				if arg.DefaultValue == nil && *arg.Hold.(*string) == stringDefaultValue {
					printProgramUsage(*env, cmd, arg.Name, arg.Description)
				}
			}
			return
		}
	}
	env.Errorf("Unknown command '%v'.\n", args[1])
	printProgramUsage(*env)
}

func InitCommandArg(env *Environment) {
	SetArgToCmd(*env, FlagRegister, FlagArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: false,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one.",
	})
	SetArgToCmd(*env, FlagIntrospect, FlagArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: true,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one. It will give more type.",
	})
}
