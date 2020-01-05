package module

import (
	"fmt"
	"github.com/EpiSchedule/src/endpoint"
	"log"
	"time"
)

func (module ModuleStruct) isUserRegistered() bool {
	for _, right := range module.Rights {
		switch right {
		case "visible":
		case "assistant":
			return true
		default:
			return module.Registered != 0
		}
	}
	return false
}

func (module ModuleStruct) isModuleStarted() bool {
	start, err := endpoint.GetDateFromString(module.BeginActive)
	if err == nil {
		return time.Now().After(start)
	}
	return false
}

func (module ModuleStruct) isModuleEnded() bool {
	end, err := endpoint.GetDateFromString(module.EndActive)
	if err == nil {
		return time.Now().After(end)
	}
	return false
}

func (module ModuleStruct) timeLeftBeforeModuleStart() string {
	if date, err := endpoint.GetDateFromString(module.BeginEvent); err != nil {
		log.Printf("An error occurred in retrieving date from string: %v\n", err.Error())
		return "null"
	} else {
		date = date.Local().Add(-time.Hour * 2) // Epitech's intranet is UTC+2
		reachTime := time.Until(date)
		return fmt.Sprintf("%v day(s) - %v hour(s) - %v minut(s)", int(reachTime.Hours()/24),
			int(reachTime.Hours())%24, int(reachTime.Minutes())%60)
	}
}
