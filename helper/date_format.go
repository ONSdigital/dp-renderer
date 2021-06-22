package helper

import (
	"html/template"
	"time"

	"github.com/ONSdigital/log.go/log"
)

func DateFormat(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Event(nil, "failed to parse time", log.Error(err), log.ERROR)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("02 January 2006"))
}

func DateFormatYYYYMMDD(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Event(nil, "failed to parse time", log.Error(err), log.ERROR)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("2006/01/02"))
}

func localiseTime(t *time.Time) time.Time {
	tz, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Event(nil, "failed to load time zone location", log.Error(err), log.ERROR)
		return *t
	}
	return t.In(tz)
}
