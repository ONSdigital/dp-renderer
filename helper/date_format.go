package helper

import (
	"context"
	"html/template"
	"regexp"
	"strings"
	"time"

	"github.com/ONSdigital/log.go/v2/log"
)

var tz *time.Location

func dateWithFormat(s string, format string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTMLEscapeString(s)
	}
	t = localiseTime(&t)
	return template.HTMLEscapeString(t.Format(format))
}

func DateFormat(s string) string {
	return dateWithFormat(s, "02 January 2006")
}

// TimeFormat24h extracts time value for 24hr clock from ISO8601 formatted timestamps
func TimeFormat24h(s string) string {
	return dateWithFormat(s, "15:04")
}

// TimeForma12h extracts time value for 12hr clock from ISO8601 formatted timestamps
func TimeFormat12h(s string) string {
	return dateWithFormat(s, "03:04pm")
}

func DateTimeFormat(s string) string {
	return dateWithFormat(s, "02 January 2006 15:04")
}

func DateFormatYYYYMMDD(s string) string {
	return dateWithFormat(s, "2006/01/02")
}

func DateFormatYYYYMMDDHyphenated(s string) string {
	return dateWithFormat(s, "2006-01-02")
}

func DateFormatYYYYMMDDNoSlash(s string) string {
	return dateWithFormat(s, "20060102")
}

func DateTimeOnsDatePatternFormat(s, lang string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTMLEscapeString(s)
	}
	t = localiseTime(&t)
	formattedTimestamp := t.Format("2 January 2006 3:04pm")

	months := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	twelveHours := []string{
		"am",
		"pm",
	}

	localeReplace := func(phrase, lang, keyRoot string, keySuffixes []string) string {
		re, _ := regexp.Compile(strings.Join(keySuffixes, "|"))
		replacer := func(match []byte) []byte {
			return []byte(Localise(keyRoot+string(match), lang, 1))
		}
		return string(re.ReplaceAllFunc([]byte(phrase), replacer))
	}

	formattedTimestamp = localeReplace(formattedTimestamp, lang, "TimestampMonth", months)
	formattedTimestamp = localeReplace(formattedTimestamp, lang, "TimestampTwelveHour", twelveHours)

	return template.HTMLEscapeString(formattedTimestamp)
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
