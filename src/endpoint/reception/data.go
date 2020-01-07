package reception

import (
	"encoding/json"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/endpoint"
	"github.com/Dayrion/EpiSchedule/src/environment"
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

func GetCurrentUserSemesterAndCredits(env environment.Environment) (int, blueprint.Credits) {
	var userReception blueprint.Reception
	if response, err := http.Get(endpoint.EpitechStartPoint + env.GetAuthentication() + blueprint.ReceptionEndpoint); err != nil {
		log.Fatal("Invalid response: " + err.Error())
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		log.Fatal("Invalid read: " + err.Error())
	} else if err := json.Unmarshal(body, &userReception); err != nil {
		log.Fatal("Invalid unmarshal: " + err.Error())
	}
	if len(userReception.Current) == 0 {
		return 0, blueprint.Credits{
			Minimum:   0,
			Aimed:     30,
			Objective: 60,
		}
	}
	return getIntFromStr(userReception.Current[0].SemesterNum), blueprint.Credits{
		Minimum:   getIntFromStr(userReception.Current[0].CreditsMin),
		Aimed:     getIntFromStr(userReception.Current[0].CreditsNorm),
		Objective: getIntFromStr(userReception.Current[0].CreditsObj),
	}
}