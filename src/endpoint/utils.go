package endpoint

import (
	"log"
	"time"
)

const (
	EpitechStartPoint = "https://intra.epitech.eu/"
)

func GetDateFromString(strDate string) (time.Time, error) {
	const layout = "2006-01-02 15:04:05"
	if date, err := time.Parse(layout, strDate); err != nil {
		return time.Time{}, err
	} else {
		return date, err
	}
}

func IsDateAfterNow(date string) bool {
	if dateTime, err := GetDateFromString(date); err != nil {
		log.Fatalf("Unable to retrieve date from string '%v'. Err: %v\n", date, err.Error())
	} else {
		return time.Now().After(dateTime)
	}
	return false
}

func IsDateBeforeNow(date string) bool {
	return !IsDateAfterNow(date)
}

/*
func RetrieveEndPointInfo(user connection.User, url string, blueprint []interface{}) (error, interface{}) {
	fmt.Print(EpitechStartPoint + user.GetAuthentication() + url + "\n")
	if response, err := http.Get(EpitechStartPoint + user.GetAuthentication() + url); err != nil {
		return err, nil
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		return err, nil
	} else {
		err := json.Unmarshal(body, &blueprint)
		return err, blueprint
	}
}*/
