package reception

import (
	"encoding/json"
	"github.com/cruffinoni/EpiSchedule/src/blueprint"
	"github.com/cruffinoni/EpiSchedule/src/environment"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getIntFromStr(number string) int {
	if nb, err := strconv.Atoi(number); err != nil {
		log.Fatalf("Error un atoi function: %v\n", err.Error())
	} else {
		return nb
	}
	return 0
}

func trySecondReceptionFormat(body []byte) (int, blueprint.Credits) {
	var userReception blueprint.ReceptionSingle
	if err := json.Unmarshal(body, &userReception); err != nil {
		log.Printf("(Body): '%v'\n", string(body))
		log.Fatal("Invalid unmarshal for both receptions: " + err.Error())
	}
	return getIntFromStr(userReception.Current.SemesterNum), blueprint.Credits{
		Minimum:   getIntFromStr(userReception.Current.CreditsMin),
		Aimed:     getIntFromStr(userReception.Current.CreditsNorm),
		Objective: getIntFromStr(userReception.Current.CreditsObj),
	}
}

func GetCurrentUserSemesterAndCredits(env environment.Environment) (int, blueprint.Credits) {
	var userReception blueprint.Reception
	if response, err := http.Get(blueprint.EpitechStartPoint + env.GetAuthentication() + blueprint.ReceptionEndpoint); err != nil {
		log.Fatal("Invalid response: " + err.Error())
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		log.Fatal("Invalid read: " + err.Error())
	} else if err := json.Unmarshal(body, &userReception); err != nil {
		return trySecondReceptionFormat(body)
	}
	return getIntFromStr(userReception.Current[0].SemesterNum), blueprint.Credits{
		Minimum:   getIntFromStr(userReception.Current[0].CreditsMin),
		Aimed:     getIntFromStr(userReception.Current[0].CreditsNorm),
		Objective: getIntFromStr(userReception.Current[0].CreditsObj),
	}
}
