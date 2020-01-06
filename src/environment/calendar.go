package environment

import (
	"blueprint"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	EpiScheduleCalendarName = "Epitech Schedule"
	defaultTimeZone         = "Europe/Paris"
)

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Retrieve a token, saves the token, then returns the generated client.
func refreshToken(config *oauth2.Config) *oauth2.Token {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := getGoogleFolder() + "token.json"
	token, err := tokenFromFile(tokFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokFile, token)
	}
	return token
}

func getGoogleFolder() string {
	switch osName := runtime.GOOS; osName {
	case "windows":
		{
			if home := os.Getenv("USERPROFILE"); home == "" {
				log.Fatal("Unable to retrieve home path from env variable.\n")
			} else {
				return home + "\\.google\\"
			}
		}
	case "linux":
		{
			if home := os.Getenv("HOME"); home == "" {
				log.Fatal("Unable to retrieve home path from env variable.\n")
			} else {
				return home + "/.google/"
			}
		}
	default:
		log.Fatal("Unknown OS or not supported by EpiShedule\n")
	}
	return ""
}

func (env *Environment) createCalendarService() {
	env.googleCalendar = &GoogleCalendar{
		registeredEvents: make(map[string]int),
	}
	credFile, err := ioutil.ReadFile(getGoogleFolder() + "credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(credFile, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	token := refreshToken(config)
	calendarService, err := calendar.NewService(env.ctx, option.WithTokenSource(config.TokenSource(env.ctx, token)))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	env.Log(VerboseDebug, "Google calendar service created from user's token.\n")
	env.googleCalendar.service = calendarService
}

func (env Environment) createEpiScheduleCalendar() *calendar.Calendar {
	createdCalendar, err := env.googleCalendar.service.Calendars.Insert(&calendar.Calendar{
		ConferenceProperties: nil,
		Summary:              EpiScheduleCalendarName,
		Description:          "Generated calendar from EpiSchedule",
		Kind:                 "calendar#calendarListEntry",
		TimeZone:             "Europe/Paris",
	}).Do()
	if err != nil {
		log.Fatalf("Unable to create a calendar: %v\n", err.Error())
	}
	env.googleCalendar.service.Calendars.Update(createdCalendar.Id, &calendar.Calendar{})
	return createdCalendar
}

func (env *Environment) retrieveCalendar() {
	calendarListService := env.googleCalendar.service.CalendarList.List()
	list, err := calendarListService.Do()
	if err != nil {
		log.Fatal("Unable to do calendar list. EpiSchedule has no right to list your calendars\n")
	}
	marshalCalendar, err := list.MarshalJSON()
	if err != nil {
		log.Fatalf("Unable to marshal calendar list: %v\n", err.Error())
	}
	var calendarList blueprint.CalendarList
	if err = json.Unmarshal(marshalCalendar, &calendarList); err != nil {
		log.Fatalf("Unable to unmarshal calendar list: %v\n", err.Error())
	}
	for _, retrievedCalendar := range calendarList.Items {
		if retrievedCalendar.Summary == EpiScheduleCalendarName {
			env.Log(VerboseDebug, "EpiSchedule calendar found.\n")
			env.googleCalendar.internalCalendar, err = env.googleCalendar.service.Calendars.Get(retrievedCalendar.ID).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve calendar's data (id: %v), error: %v\n",
					retrievedCalendar.ID, err.Error())
			}
			return
		}
	}
	env.Log(VerboseDebug, "No calendar generated by EpiSchedule found, generating a new one.\n")
	env.googleCalendar.internalCalendar = env.createEpiScheduleCalendar()
}

func (env *Environment) listRegisteredEvents() {
	events, err := env.googleCalendar.service.Events.List(env.googleCalendar.internalCalendar.Id).Do()
	if err != nil {
		log.Fatalf("Unable to make a list of all events generated in EpiSchedule calendar: %v\n", err.Error())
	}
	marshalEvents, err := events.MarshalJSON()
	if err != nil {
		log.Fatalf("Unable to marshal calendar list: %v\n", err.Error())
	}
	var eventsList blueprint.CalendarEventList
	if err = json.Unmarshal(marshalEvents, &eventsList); err != nil {
		log.Fatalf("Unable to unmarshal calendar list: %v\n", err.Error())
	}
	for _, event := range eventsList.Items {
		if event.Summary == "" {
			continue
		}
		env.googleCalendar.registeredEvents[event.Summary]++
	}
}

func dateToRFC3339(strDate string) string {
	const layout = "2006-01-02 15:04:05 -0700 UTC"
	strDate = strDate + " +0100 UTC"
	if date, err := time.Parse(layout, strDate); err != nil {
		log.Fatal("Unable to parse a date.\n")
	} else {
		return date.Format(time.RFC3339)
	}
	return ""
}

func (env Environment) AddEvent(title, description, typeTitle, begin, end string) {
	if env.googleCalendar.registeredEvents[title] > 0 {
		env.Logf(VerboseSimple, "	> The event '%v' is already registered in the calendar.\n", title)
		return
	}
	_, err := env.googleCalendar.service.Events.Insert(env.googleCalendar.internalCalendar.Id,
		&calendar.Event{
			AnyoneCanAddSelf: false,
			ColorId:          activityColor[typeTitle],
			Description:      description,
			End: &calendar.EventDateTime{
				DateTime: dateToRFC3339(end),
				TimeZone: defaultTimeZone,
			},
			Kind:     "calendar#event",
			Location: "Epitech, Strasbourg, France",
			Start: &calendar.EventDateTime{
				DateTime: dateToRFC3339(begin),
				TimeZone: defaultTimeZone,
			},
			Status:       "confirmed",
			Summary:      title,
			Transparency: "opaque",
			Visibility:   "private",
		}).Do()
	if err != nil {
		log.Fatalf("Unable to insert an event: '%v'\n", err.Error())
	}
	env.googleCalendar.registeredEvents[title]++
	env.Log(VerboseSimple, "	> Activity successfully added.\n")

}
