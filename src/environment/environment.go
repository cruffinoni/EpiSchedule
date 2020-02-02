package environment

import (
	"context"
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"log"
	"net/http"
)

const (
	ProjectName = "EpiSchedule"
)

type UserData struct {
	Semester int
	Credits  blueprint.Credits
}

type Environment struct {
	authentication  string
	Client          *http.Client
	User            UserData
	verbose         VerboseLevel
	autoRegister    []string
	autoAddCalendar []string
	googleCalendar  *GoogleCalendar
	ctx             context.Context
	Flag            Flag
}

func NewEnvironment() Environment {
	fmt.Print("Program initialization...\n")
	env := Environment{
		authentication: GetAuthLoginLinkFromEnv(),
		Client:         http.DefaultClient,
		User: UserData{
			Semester: 0,
		},
		verbose:        VerboseDefault,
		googleCalendar: nil,
		ctx:            context.Background(),
	}
	if env.authentication == "" {
		log.Fatal("Unable to retrieve autologin link from env\n")
	}
	env.Log(VerboseDebug, "Testing autologin link.\n")
	testConnection(env)
	env.Log(VerboseDebug, "Initialization successful.\n")
	return env
}

func (env *Environment) SetUpCalendar() {
	if env.googleCalendar != nil {
		return
	}
	env.Logf(VerboseDebug, "Creating %v Google Calendar if needed.\n", ProjectName)
	env.createCalendarService()
	env.Logf(VerboseDebug, "Retrieve registered activities in %v's Google Calendar.\n", ProjectName)
	env.retrieveCalendar()
	env.listRegisteredEvents()
}
