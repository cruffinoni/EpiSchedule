package flag

import (
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
)

func SetArgToCmd(cmd string, arg ProgArg) {
	if _, ok := cmdArg[cmd]; !ok {
		log.Fatalf("nonexistent cmd: %v\n", cmd)
	}
	if !cmdArg[cmd].Args.IsEqual(ArgEmpty) {
		log.Fatalf("cmd %v has already args set.\n", cmd)
	}
	cmdArg[cmd].Args = arg
}

func SetPreHandlerToCmd(cmd string, handler func()) {
	if _, ok := cmdArg[cmd]; !ok {
		log.Fatalf("nonexistent cmd: %v\n", cmd)
	}
	cmdArg[cmd].preHandler = handler
}

func SetHandlerToCmd(cmd string, handlerCmd HandlerCmd) {
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
}
