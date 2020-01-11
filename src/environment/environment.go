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
	testConnection(env)
	fmt.Print("Initialization done...\n")
	return env
}

func (env *Environment) SetUpCalendar() {
	if env.googleCalendar != nil {
		return
	}
	env.createCalendarService()
	env.retrieveCalendar()
	env.listRegisteredEvents()
}
