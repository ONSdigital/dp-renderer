package helper

import (
	"strconv"
	"strings"
	"time"
)

// DatePeriodFormat will format a time-series date period string to a human accessible format e.g.
// "2019 JAN-FEB" to "Jan - Feb 2019"
// "2010 Q1" to "Jan - Mar 2010"
func DatePeriodFormat(s string) string {
	dashIndex := strings.Index(s, "-")
	// 1. Add spaces around dash
	if dashIndex > -1 {
		charIndexAfterDash := dashIndex + 1
		s = s[:dashIndex] + " - " + s[charIndexAfterDash:]
	}

	// 2. Replace Q1 Q2 Q3 Q4 with their quarterly month representation e.g. Apr - Jun
	Q1 := strings.Index(s, "Q1")
	if Q1 > -1 {
		Q1EndIndex := Q1 + 2
		s = s[:Q1] + "Jan - Mar" + s[Q1EndIndex:]
	}
	Q2 := strings.Index(s, "Q2")
	if Q2 > -1 {
		Q2EndIndex := Q2 + 2
		s = s[:Q2] + "Apr - Jun" + s[Q2EndIndex:]
	}
	Q3 := strings.Index(s, "Q3")
	if Q3 > -1 {
		Q3EndIndex := Q3 + 2
		s = s[:Q3] + "Jul - Sep" + s[Q3EndIndex:]

	}
	Q4 := strings.Index(s, "Q4")
	if Q4 > -1 {
		Q4EndIndex := Q4 + 2
		s = s[:Q4] + "Oct - Dec" + s[Q4EndIndex:]

	}

	// 3. Move year to end of string if present and insert a space
	year, err := strconv.Atoi(s[:4])
	if err == nil {
		// Not just displaying year but month as well
		// YYYY[space]DEC[space]-[space]JAN = 14 characters
		if len(s) == 14 {
			monthStart, _ := time.Parse("Jan", s[5:8])
			monthEnd, _ := time.Parse("Jan", s[11:])

			dateStart := monthStart.AddDate(year, 0, 0)
			dateEnd := monthEnd.AddDate(year, 0, 0)

			if dateStart.After(dateEnd) {
				dateEnd = dateEnd.AddDate(1, 0, 0)
				s = dateStart.Format("Jan 2006")
			} else {
				s = dateStart.Format("Jan")
			}
			s = s + " - " + dateEnd.Format("Jan 2006")
		} else if len(s) > 5 {
			// YYYY[space] = 5 characters
			postYearIndex := 5
			s = s[postYearIndex:] + " " + s[:4]
		}
	}

	// 4. Convert BLOCK CAPS to Title Caps
	timePeriodFormatted := strings.Title(strings.ToLower(s))
	return timePeriodFormatted
}
