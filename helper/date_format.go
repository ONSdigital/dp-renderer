package helper

import (
	"context"
	"html/template"
	"time"

	"github.com/ONSdigital/log.go/v2/log"
)

var tz *time.Location

func DateFormat(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTMLEscapeString(s)
	}
	t = localiseTime(&t)
	return template.HTMLEscapeString(t.Format("02 January 2006"))
}

func DateTimeFormat(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTML(s)
	}
	t = localiseTime(&t)
	return template.HTML(t.Format("02 January 2006 15:04"))
}

func DateFormatYYYYMMDD(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTMLEscapeString(s)
	}
	t = localiseTime(&t)
	return template.HTMLEscapeString(t.Format("2006/01/02"))
}

func DateFormatYYYYMMDDNoSlash(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTMLEscapeString(s)
	}
	t = localiseTime(&t)
	return template.HTMLEscapeString(t.Format("20060102"))
}

func init() {
	var err error
	if tz, err = time.LoadLocation("Europe/London"); err != nil {
		log.Error(context.Background(), "failed to load timezone", err)
		tz = nil
	}
}

func localiseTime(t *time.Time) time.Time {
	if tz == nil {
		return *t
	}
	return t.In(tz)
}
