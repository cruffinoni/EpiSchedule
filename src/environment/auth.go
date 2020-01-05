package environment

import (
	"os"
)

func (env Environment) GetAuthentication() string {
	return env.authentication
}

func GetAuthLoginLinkFromEnv() string {
	return os.Getenv("EPISCHEDULE_AUTOLOGIN_LINK")
}
