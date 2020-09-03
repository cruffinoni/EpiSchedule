# EpiSchedule
- Automatic inscription to desired modules and activities.
- Automatic ingestion of incoming activities to your Google Calendar.

## Dependencies
```bash
$ go get -u google.golang.org/api/calendar/v3
$ go get -u golang.org/x/oauth2/google
```

### Todo:
- Add vendor for dependencies
- Setup tutorial about Google Quickstart
- Get automatically the current school year
- Don't add module nor event if it's already there but move it if the date/hours differ
- Warn the amount of credits is too low (like the intranet does)
- Pitch aren't recognized as an activity - Should I edit the whole course system detection?
- Add a commit norm for future commits
