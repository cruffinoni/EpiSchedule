package flag

import (
	"flag"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
	"os"
	"strings"
)

var (
	cmdArg = make(map[string]*ProgCmd)
)

func listCmd() []string {
	cmdList := make([]string, 0)
	for cmd := range cmdArg {
		cmdList = append(cmdList, cmd)
	}
	if len(cmdList) == 0 {
		return []string{
			"no command found, contact the developer",
		}
	} else {
		return cmdList
	}
}

func printProgramUsage(env environment.Environment, optional ...string) {
	if len(optional) < 3 {
		env.Errorf("Usage of %v:\n\t./%v [command] [arguments]\n\t\t"+
			"command: One of these main commands: %v\n\t\t"+
			"arguments: The argument(s) associated to the command\n",
			environment.ProjectName, strings.ToLower(environment.ProjectName), listCmd())
	} else {
		env.Errorf("Usage of %v:\n\t./%v %v [arguments]\n\t\t"+
			"%v: %v\n",
			environment.ProjectName, strings.ToLower(environment.ProjectName), optional[0],
			optional[1], optional[2])
	}
	os.Exit(1)
}

func RetrieveCommand(env *environment.Environment, args []string) *ProgCmd {
	if len(args) < 2 {
		printProgramUsage(*env)
	}
	for cmdName, cmd := range cmdArg {
		if args[1] == cmdName {
			argCmdSet := flag.NewFlagSet(cmdName, flag.ExitOnError)
			for _, arg := range cmd.Args {
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
			}
			_ = argCmdSet.Parse(args[2:])
			for _, arg := range cmd.Args {
				switch arg.Hold.(type) {
				case *int:
					if arg.DefaultValue == nil && *arg.Hold.(*int) == intDefaultValue {
						printProgramUsage(*env, cmdName, arg.Name, arg.Description)
					}
				case *bool:
					if arg.DefaultValue == nil && *arg.Hold.(*bool) == boolDefaultValue {
						printProgramUsage(*env, cmdName, arg.Name, arg.Description)
					}
				case *string:
					if arg.DefaultValue == nil && *arg.Hold.(*string) == stringDefaultValue {
						printProgramUsage(*env, cmdName, arg.Name, arg.Description)
					}
				}
			}
			return cmd
		}
	}
	env.Errorf("Unknown command '%v'.\n", args[1])
	printProgramUsage(*env)
	return nil
}

func (cmd *ProgCmd) ExecuteHandlers(env *environment.Environment, courses []blueprint.Course) {
	if cmd.preHandler != nil {
		cmd.preHandler(env)
	}
	if cmd.handler != nil {
		cmd.handler(*env, courses)
	}
}
