package helper

import (
	"html/template"
)

var RegisteredFuncs template.FuncMap = template.FuncMap{
	"humanSize":                    HumanSize,
	"stringArrayContains":          StringArrayContains,
	"safeHTML":                     SafeHTML,
	"timeFormat24h":                TimeFormat24h,
	"timeFormat12h":                TimeFormat12h,
	"dateFormat":                   DateFormat,
	"dateTimeFormat":               DateTimeFormat,
	"dateFormatYYYYMMDD":           DateFormatYYYYMMDD,
	"dateFormatYYYYMMDDNoSlashes":  DateFormatYYYYMMDDNoSlash,
	"dateFormatYYYYMMDDHyphenated": DateFormatYYYYMMDDHyphenated,
	"datePeriodFormat":             DatePeriodFormat,
	"dateTimeOnsDatePatternFormat": DateTimeOnsDatePatternFormat,
	"last":                         Last,
	"loop":                         Loop,
	"add":                          Add,
	"subtract":                     Subtract,
	"multiply":                     Multiply,
	"slug":                         Slug,
	"legacyDatasetDownloadURI":     LegacyDatasetDownloadURI,
	"markdown":                     Markdown,
	"localise":                     Localise,
	"domainSetLang":                DomainSetLang,
	"hasField":                     HasField,
	"notLastItem":                  NotLastItem,
	"concatenateStrings":           ConcatenateStrings,
	"truncateToMaximuCharacters":   TruncateToMaximumCharacters,
	"trimPrefixedPeriod":           TrimPrefixedPeriod,
	"intToString":                  IntToString,
	"lower":                        Lower,
	"thousandsSeparator":           ThousandsSeparator,
}
