package flag

import (
	"flag"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
	"os"
	"strings"
)

var ArgEmpty ProgArg

var (
	validCmd = ArgType{
		ArgRegister,
		ArgShow,
		ArgIntrospect,
	}

	cmdArg = map[string]*ProgCmd{
		ArgRegister:   new(ProgCmd),
		ArgShow:       new(ProgCmd),
		ArgIntrospect: new(ProgCmd),
	}
)

func printProgramUsage(env environment.Environment, optional ...string) {
	if len(optional) < 3 {
		env.Errorf("Usage of %v:\n\t./%v [command] [arguments]\n\t\t"+
			"command: One of these main commands: %v\n\t\t"+
			"arguments: The argument(s) associated to the command\n",
			environment.ProjectName, strings.ToLower(environment.ProjectName), validCmd)
	} else {
		env.Errorf("Usage of %v:\n\t./%v %v [arguments]\n\t\t"+
			"%v: %v\n",
			environment.ProjectName, strings.ToLower(environment.ProjectName), optional[0],
			optional[1], optional[2])
	}
	os.Exit(1)
}

func RetrieveCommandFlag(env *environment.Environment, args []string) *ProgCmd {
	if len(args) < 2 {
		printProgramUsage(*env)
	}
	for cmdName, cmd := range cmdArg {
		if !cmd.Args.IsEqual(ArgEmpty) && args[1] == cmdName {
			argCmdSet := flag.NewFlagSet(cmdName, flag.ExitOnError)
			switch cmd.Args.Hold.(type) {
			case nil:
				log.Fatalf("command %v has an nil type\n", cmd.Args.Name)
			case *int:
				defaultValue := intDefaultValue
				if cmd.Args.DefaultValue != nil {
					defaultValue = cmd.Args.DefaultValue.(int)
				}
				argCmdSet.IntVar(cmd.Args.Hold.(*int), cmd.Args.Name, defaultValue, cmd.Args.Description)
			case *bool:
				defaultValue := boolDefaultValue
				if cmd.Args.DefaultValue != nil {
					defaultValue = cmd.Args.DefaultValue.(bool)
				}
				argCmdSet.BoolVar(cmd.Args.Hold.(*bool), cmd.Args.Name, defaultValue, cmd.Args.Description)
			case *string:
				defaultValue := stringDefaultValue
				if cmd.Args.DefaultValue != nil {
					defaultValue = cmd.Args.DefaultValue.(string)
				}
				argCmdSet.StringVar(cmd.Args.Hold.(*string), cmd.Args.Name, defaultValue, cmd.Args.Description)
			default:
				log.Fatalf("command %v has an unknown type\n", cmd.Args.Name)
			}
			_ = argCmdSet.Parse(args[2:])
			switch cmd.Args.Hold.(type) {
			case *int:
				if cmd.Args.DefaultValue == nil && *cmd.Args.Hold.(*int) == intDefaultValue {
					printProgramUsage(*env, cmdName, cmd.Args.Name, cmd.Args.Description)
				}
			case *bool:
				if cmd.Args.DefaultValue == nil && *cmd.Args.Hold.(*bool) == boolDefaultValue {
					printProgramUsage(*env, cmdName, cmd.Args.Name, cmd.Args.Description)
				}
			case *string:
				if cmd.Args.DefaultValue == nil && *cmd.Args.Hold.(*string) == stringDefaultValue {
					printProgramUsage(*env, cmdName, cmd.Args.Name, cmd.Args.Description)
				}
			}
			return cmd
		}
	}
	env.Errorf("Unknown command '%v'.\n", args[1])
	printProgramUsage(*env)
	return nil
}

func (cmd *ProgCmd) ExecuteHandlers(env environment.Environment, courses []blueprint.Course) {
	if cmd.preHandler != nil {
		cmd.preHandler()
	}
	if cmd.handler != nil {
		cmd.handler(env, courses)
	}
}
