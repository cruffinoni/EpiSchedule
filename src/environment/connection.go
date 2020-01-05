package environment

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testConnection(user Environment) {
	res, err := http.Get(fmt.Sprintf("https://intra.epitech.eu/%v/planning/load?format=json&start=2019-10-07&end=2019-10-13", user.GetAuthentication()))
	if err != nil {
		log.Fatalf("Something wrong happened. Got the error '%v'\n", err.Error())
	}
	if _, err := ioutil.ReadAll(res.Body); err != nil {
		log.Fatalf("Unable to read the request result's body")
	}
}
