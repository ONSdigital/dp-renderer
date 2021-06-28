package helper

import (
	"html/template"
	"time"

	"github.com/ONSdigital/log.go/v2/log"
)

func DateFormat(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(nil, "failed to parse time", err)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("02 January 2006"))
}

func DateFormatYYYYMMDD(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(nil, "failed to parse time", err)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("2006/01/02"))
}

func localiseTime(t *time.Time) time.Time {
	tz, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Error(nil, "failed to load timezone", err)
		return *t
	}
	return t.In(tz)
}
