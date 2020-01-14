package flag

import (
	"log"
)

func SetArgToCmd(cmd string, arg ProgArg) {
	if _, ok := cmdArg[cmd]; !ok {
		cmdArg[cmd] = new(ProgCmd)
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
		cmdArg[cmd] = new(ProgCmd)
	}
	cmdArg[cmd].preHandler = handler
}

func SetHandlerToCmd(cmd string, handlerCmd HandlerType) {
	if _, ok := cmdArg[cmd]; !ok {
		cmdArg[cmd] = new(ProgCmd)
	}
	cmdArg[cmd].handler = handlerCmd
}
