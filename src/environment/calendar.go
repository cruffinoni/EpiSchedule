package environment

import (
	"encoding/json"
	"fmt"
	"github.com/cruffinoni/EpiSchedule/src/blueprint"
	"github.com/cruffinoni/EpiSchedule/src/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	EpiScheduleCalendarName = "Epitech Schedule"
	defaultTimeZone         = "Europe/Paris"
	googleFolder            = "./google/"
)

type GoogleCalendar struct {
	service          *calendar.Service
	internalCalendar *calendar.Calendar
	registeredEvents map[string]blueprint.CalendarEvent
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(env Environment, config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	env.Log(VerboseSimple, ColorYellow+"! To access your to calendar, "+ProjectName+" needs your authorization. "+
		"A link to Google will open allowing you to connect, securely, to it.\n"+
		"Once connected, enter the authorization code: \n")
	time.Sleep(time.Second * 5)
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", authURL).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", authURL).Start()
	case "darwin":
		err = exec.Command("open", authURL).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(env.ctx, authCode)
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
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	defer f.Close()
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
func refreshToken(env Environment, config *oauth2.Config) *oauth2.Token {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := googleFolder + "token.json"
	token, err := tokenFromFile(tokFile)
	if err != nil {
		token = getTokenFromWeb(env, config)
		saveToken(tokFile, token)
	}
	return token
}

func (env *Environment) createCalendarService() {
	env.googleCalendar = &GoogleCalendar{
		registeredEvents: make(map[string]blueprint.CalendarEvent),
	}
	credFile, err := ioutil.ReadFile(googleFolder + "credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(credFile, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	token := refreshToken(*env, config)
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
		log.Fatalf("Unable to do calendar list. EpiSchedule has no right to list your calendars: %v\n", err.Error())
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
		env.googleCalendar.registeredEvents[event.Summary] = event
	}
}

func (env Environment) RemoveEvent(title string) {
	if event, ok := env.googleCalendar.registeredEvents[title]; !ok {
		env.Logf(VerboseDebug, "Impossible to remove: '%v'. The event doesn't exists.\n", title)
		return
	} else {
		r := env.googleCalendar.service.Events.Delete(env.googleCalendar.internalCalendar.Id,
			event.ID)
		err := r.Do()
		if err == nil {
			env.Errorf("Unable to delete the event: %v - Bad return from Google\n", title)
		}
	}
}

func (env Environment) AddEvent(activity blueprint.CourseActivity) {
	if e, ok := env.googleCalendar.registeredEvents[activity.Title]; ok {
		startDate, err := utils.GetCESTDateFromString(activity.Events[0].Begin)
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Printf("Date are not even: %+v & %+v = %v\n", startDate, e.Start.DateTime, activity.Events[0].Begin)
		if utils.IsDateEven(startDate, e.Start.DateTime) {
			env.Logf(VerboseSimple, ColorMagenta+"	> The event '%v' is already registered in the calendar.\n", activity.Title)
			return
		} else {
			env.Logf(VerboseSimple, ColorMagenta+"	> The event '%v' is already registered but the date has changed.\n", activity.Title)
			//env.RemoveEvent(activity.Title)
			//env.Log(VerboseSimple, ColorMagenta+"	> Event deleted from the calendar. A new will be added with the updated hours.\n")
		}
		return
	}
	extractedLocation := strings.Split(activity.Events[0].Location, "/")
	locationName := ""
	if extractedLocation[len(extractedLocation)-1] != "" {
		locationName = strings.Replace(extractedLocation[len(extractedLocation)-1], "-", " ", -1)
	} else {
		locationName = "N/A"
	}
	color := AvailableActivity[activity.TypeTitle]
	if AvailableActivity[activity.TypeTitle] == "" {
		color = "1"
		env.Errorf("Activity named '%v' is unknown to %v. Back to the default color.\n", activity.TypeTitle, ProjectName)
	}
	_, err := env.googleCalendar.service.Events.Insert(env.googleCalendar.internalCalendar.Id,
		&calendar.Event{
			AnyoneCanAddSelf: false,
			ColorId:          color,
			Description:      activity.Description,
			End: &calendar.EventDateTime{
				DateTime: utils.FullDateToRFC3339(activity.Events[0].End),
				TimeZone: defaultTimeZone,
			},
			Kind:     "calendar#event",
			Location: locationName,
			Start: &calendar.EventDateTime{
				DateTime: utils.FullDateToRFC3339(activity.Events[0].Begin),
				TimeZone: defaultTimeZone,
			},
			Status:       "confirmed",
			Summary:      activity.Title,
			Transparency: "opaque",
			Visibility:   "private",
		}).Do()
	if err != nil {
		log.Fatalf("Unable to insert an event: '%v'\n", err.Error())
	}
	env.googleCalendar.registeredEvents[activity.Title] = blueprint.CalendarEvent{
		Start: struct {
			DateTime time.Time `json:"dateTime"`
		}{
			DateTime: time.Now(),
		},
		Summary: activity.Title,
	}
	env.Log(VerboseSimple, ColorBlue+"	> Activity successfully added.\n")
}

func (env Environment) AddModule(module blueprint.CourseSummary, project blueprint.CourseActivity) {
	if _, ok := env.googleCalendar.registeredEvents[project.Title]; ok {
		env.Logf(VerboseSimple, ColorCyan+"Module %v is already there\n", project.Title)
		return
	}
	_, err := env.googleCalendar.service.Events.Insert(env.googleCalendar.internalCalendar.Id,
		&calendar.Event{
			AnyoneCanAddSelf: false,
			ColorId:          fmt.Sprintf("%v", rand.Intn(11)+1),
			End: &calendar.EventDateTime{
				DateTime: utils.FullDateToRFC3339(project.End),
				TimeZone: defaultTimeZone,
			},
			Kind: "calendar#event",
			Start: &calendar.EventDateTime{
				DateTime: utils.FullDateToRFC3339(project.Begin),
				TimeZone: defaultTimeZone,
			},
			Status:       "confirmed",
			Summary:      project.Title,
			Transparency: "opaque",
			Visibility:   "private",
			Description:  fmt.Sprintf("Module: %v\n%v", module.Title, project.Description),
		}).Do()
	if err != nil {
		log.Fatalf("Unable to insert an event: '%v'\n", err.Error())
	}
}
