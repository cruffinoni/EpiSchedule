package environment

import (
	"log"
	"net/http"
)

type UserData struct {
	Semester int
}

type Environment struct {
	authentication string
	Client         *http.Client
	User           UserData
	verbose        VerboseLevel
	autoRegister   []string
}

func NewEnvironment(semester int) Environment {
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
		verbose: VerboseDefault,
	}
	testConnection(env)
	return env
}
