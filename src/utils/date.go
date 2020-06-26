package utils

import (
	"log"
	"time"
)

func GetDateFromString(strDate string) (time.Time, error) {
	const layout = "2006-01-02 15:04:05"
	if date, err := time.Parse(layout, strDate); err != nil {
		return time.Time{}, err
	} else {
		return date, err
	}
}

func GetCESTDateFromString(strDate string) (time.Time, error) {
	const layout = "2006-01-02 15:04:05 -0700 UTC"
	strDate = strDate + " +0100 UTC"
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

func DateToRFC3339(strDate string) string {
	const layout = "2006-01-02 15:04:05 -0700 UTC"
	strDate = strDate + " +0100 UTC"
	if date, err := time.Parse(layout, strDate); err != nil {
		log.Fatal("Unable to parse a date.\n")
	} else {
		return date.Format(time.RFC3339)
	}
	return ""
}

func IsDateEven(a time.Time, b time.Time) bool {
	y, m, d := a.Date()
	yy, mm, dd := b.Date()
	//log.Printf("a: %v & %v & %v\n", y, m, d)
	//log.Printf("b: %v & %v & %v\n", yy, mm, dd)
	return y == yy && m == mm && d == dd && a.Hour() == b.Hour()
}
