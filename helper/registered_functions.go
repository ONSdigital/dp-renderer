package helper

import "html/template"

var RegisteredFuncs template.FuncMap = template.FuncMap{
	"humanSize":                  HumanSize,
	"safeHTML":                   SafeHTML,
	"dateFormat":                 DateFormat,
	"dateFormatYYYYMMDD":         DateFormatYYYYMMDD,
	"datePeriodFormat":           DatePeriodFormat,
	"last":                       Last,
	"loop":                       Loop,
	"subtract":                   Subtract,
	"slug":                       Slug,
	"legacyDatasetDownloadURI":   LegacyDatasetDownloadURI,
	"markdown":                   Markdown,
	"localise":                   Localise,
	"domainSetLang":              DomainSetLang,
	"hasField":                   HasField,
	"notLastItem":                NotLastItem,
	"concatenateStrings":         ConcatenateStrings,
	"truncateToMaximuCharacters": TruncateToMaximumCharacters,
}
