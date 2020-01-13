package flag

import (
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
)

func SetArgToCmd(cmd string, arg ProgArg) {
	if _, ok := cmdArg[cmd]; !ok {
		log.Fatalf("nonexistent cmd: %v\n", cmd)
	}
	for _, cmdArg := range cmdArg[cmd].Args {
		if cmdArg.IsEqual(arg) {
			log.Fatalf("cmd %v has multiple definition of the arg: %v.\n",
				cmd, arg.Name)
		}
	}
	cmdArg[cmd].Args = append(cmdArg[cmd].Args, arg)
}

func SetUpPreHandler(cmd string, handler PreHandlerType) {
	if _, ok := cmdArg[cmd]; !ok {
		log.Fatalf("nonexistent cmd: %v\n", cmd)
	}
	cmdArg[cmd].preHandler = handler
}

func SetHandlerToCmd(cmd string, handlerCmd HandlerType) {
	if _, ok := cmdArg[cmd]; !ok {
		log.Fatalf("nonexistent cmd: %v\n", cmd)
	}
	cmdArg[cmd].handler = handlerCmd
}

func InitCommandArg(env *environment.Environment) {
	SetArgToCmd(ArgRegister, ProgArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: false,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one.",
	})
	SetArgToCmd(ArgIntrospect, ProgArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: true,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one. It will give more type.",
	})
	SetArgToCmd(ArgIntrospect, ProgArg{
		Hold:         &env.Flag.SaveActivities,
		DefaultValue: true,
		Name:         "save",
		Description:  "(Optional) Save the displayed activities.",
	})
}
