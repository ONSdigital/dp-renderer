package helper

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func ThousandsSeparator(i int) string {
	p := message.NewPrinter(language.English)
	return p.Sprint(number.Decimal(i))
}
