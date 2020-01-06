package environment

import (
	"context"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"log"
	"net/http"
)

type UserData struct {
	Semester int
}

type GoogleCalendar struct {
	service          *calendar.Service
	internalCalendar *calendar.Calendar
	registeredEvents map[string]int
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
}

func NewEnvironment(semester int) Environment {
	fmt.Print("Program initialization...\n")
	authentication := GetAuthLoginLinkFromEnv()
	if authentication == "" {
		log.Fatal("Unable to retrieve autologin link from env\n")
	}
	env := Environment{
		authentication: authentication,
		Client:         http.DefaultClient,
		User: UserData{
			Semester: semester,
		},
		verbose:        VerboseDefault,
		googleCalendar: nil,
		ctx:            context.Background(),
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
