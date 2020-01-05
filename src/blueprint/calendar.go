package blueprint

import "time"

type CalendarList struct {
	Etag  string `json:"etag"`
	Items []struct {
		AccessRole           string `json:"accessRole"`
		BackgroundColor      string `json:"backgroundColor"`
		ColorID              string `json:"colorId"`
		ConferenceProperties struct {
			AllowedConferenceSolutionTypes []string `json:"allowedConferenceSolutionTypes"`
		} `json:"conferenceProperties"`
		Etag             string `json:"etag"`
		ForegroundColor  string `json:"foregroundColor"`
		ID               string `json:"id"`
		Kind             string `json:"kind"`
		Selected         bool   `json:"selected"`
		Summary          string `json:"summary"`
		TimeZone         string `json:"timeZone"`
		DefaultReminders []struct {
			Method  string `json:"method"`
			Minutes int    `json:"minutes"`
		} `json:"defaultReminders,omitempty"`
		Primary         bool   `json:"primary,omitempty"`
		Description     string `json:"description,omitempty"`
		SummaryOverride string `json:"summaryOverride,omitempty"`
	} `json:"items"`
	Kind          string `json:"kind"`
	NextSyncToken string `json:"nextSyncToken"`
}

type CalendarEventList struct {
	AccessRole  string `json:"accessRole"`
	Description string `json:"description"`
	Etag        string `json:"etag"`
	Items       []struct {
		Created time.Time `json:"created"`
		Creator struct {
			Email string `json:"email"`
		} `json:"creator"`
		Description string `json:"description,omitempty"`
		End         struct {
			DateTime time.Time `json:"dateTime"`
		} `json:"end"`
		Etag      string `json:"etag"`
		HTMLLink  string `json:"htmlLink"`
		ICalUID   string `json:"iCalUID"`
		ID        string `json:"id"`
		Kind      string `json:"kind"`
		Location  string `json:"location,omitempty"`
		Organizer struct {
			DisplayName string `json:"displayName"`
			Email       string `json:"email"`
			Self        bool   `json:"self"`
		} `json:"organizer"`
		Reminders struct {
			UseDefault bool `json:"useDefault"`
		} `json:"reminders"`
		Start struct {
			DateTime time.Time `json:"dateTime"`
		} `json:"start"`
		Status  string    `json:"status"`
		Summary string    `json:"summary,omitempty"`
		Updated time.Time `json:"updated"`
	} `json:"items"`
	Kind          string    `json:"kind"`
	NextSyncToken string    `json:"nextSyncToken"`
	Summary       string    `json:"summary"`
	TimeZone      string    `json:"timeZone"`
	Updated       time.Time `json:"updated"`
}
