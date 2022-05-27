package helper_test

import (
	"strings"
	"testing"

	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDateFormat(t *testing.T) {
	Convey("Date format returns human readable string", t, func() {
		So(helper.DateFormat("2019-08-15T00:00:00.000Z"), ShouldEqual, "15 August 2019")
		So(helper.DateFormat("2019-05-21T23:00:00.000Z"), ShouldEqual, "22 May 2019") // BST
		So(helper.DateFormat("2019-12-21T23:00:00.000Z"), ShouldEqual, "21 December 2019")
		So(helper.DateFormat("2019-08-15"), ShouldEqual, "2019-08-15")
		So(helper.DateFormat(""), ShouldEqual, "")
	})
}

func TestDateFormatYYYYMMDD(t *testing.T) {
	Convey("Date format returns short date pattern without slashes", t, func() {
		So(helper.DateFormatYYYYMMDD("2019-08-15T00:00:00.000Z"), ShouldEqual, "2019/08/15")
		So(helper.DateFormatYYYYMMDD("2019-05-21T23:00:00.000Z"), ShouldEqual, "2019/05/22") // BST
		So(helper.DateFormatYYYYMMDD("2019-12-21T23:00:00.000Z"), ShouldEqual, "2019/12/21")
		So(helper.DateFormatYYYYMMDD("2019-08-15"), ShouldEqual, "2019-08-15")
		So(helper.DateFormatYYYYMMDD(""), ShouldEqual, "")
	})
}

func TestDateFormatYYYYMMDDNoSlash(t *testing.T) {
	Convey("Date format returns human readable string", t, func() {
		So(helper.DateFormatYYYYMMDDNoSlash("2019-08-15T00:00:00.000Z"), ShouldEqual, "20190815")
		So(helper.DateFormatYYYYMMDDNoSlash("2019-05-21T23:00:00.000Z"), ShouldEqual, "20190522") // BST
		So(helper.DateFormatYYYYMMDDNoSlash("2019-12-21T23:00:00.000Z"), ShouldEqual, "20191221")
		So(helper.DateFormatYYYYMMDDNoSlash("2019-08-15"), ShouldEqual, "2019-08-15") // failed to parse, so returns arg value
		So(helper.DateFormatYYYYMMDDNoSlash(""), ShouldEqual, "")
	})
}

var cyLocale = []string{
	"[TimestampMonthMay]",
	"one = \"Mai\"",
	"[TimestampMonthAugust]",
	"one = \"Awst\"",
	"[TimestampMonthDecember]",
	"one = \"Rhagfyr\"",
	"[TimestampTwelveHouram]",
	"one = \"am\"",
	"[TimestampTwelveHourpm]",
	"one = \"pm\"",
}

var enLocale = []string{
	"[TimestampMonthMay]",
	"one = \"May\"",
	"[TimestampMonthAugust]",
	"one = \"August\"",
	"[TimestampMonthDecember]",
	"one = \"December\"",
	"[TimestampTwelveHouram]",
	"one = \"am\"",
	"[TimestampTwelveHourpm]",
	"one = \"pm\"",
}

func mockTimestampAssetFunction(name string) ([]byte, error) {
	if strings.Contains(name, ".cy.toml") {
		return []byte(strings.Join(cyLocale, "\n")), nil
	}
	return []byte(strings.Join(enLocale, "\n")), nil
}

func TestDateTimeOnsDatePatternFormat(t *testing.T) {
	helper.InitialiseLocalisationsHelper(mockTimestampAssetFunction)

	Convey("Date format returns human readable string", t, func() {
		So(helper.DateTimeOnsDatePatternFormat("2019-08-03T00:00:00.000Z", ""), ShouldEqual, "3 August 2019 1:00am")  // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-08-15T00:00:00.000Z", ""), ShouldEqual, "15 August 2019 1:00am") // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-05-21T23:00:00.000Z", ""), ShouldEqual, "22 May 2019 12:00am")   // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-12-21T23:00:00.000Z", ""), ShouldEqual, "21 December 2019 11:00pm")
		So(helper.DateTimeOnsDatePatternFormat("2019-08-15", ""), ShouldEqual, "2019-08-15") // failed to parse, so returns arg value
		So(helper.DateTimeOnsDatePatternFormat("", ""), ShouldEqual, "")
	})

	Convey("Date is localised", t, func() {
		So(helper.DateTimeOnsDatePatternFormat("2019-08-03T00:00:00.000Z", "en"), ShouldEqual, "3 August 2019 1:00am")  // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-08-15T00:00:00.000Z", "en"), ShouldEqual, "15 August 2019 1:00am") // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-05-21T23:00:00.000Z", "en"), ShouldEqual, "22 May 2019 12:00am")   // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-12-21T23:00:00.000Z", "en"), ShouldEqual, "21 December 2019 11:00pm")

		So(helper.DateTimeOnsDatePatternFormat("2019-08-03T00:00:00.000Z", "cy"), ShouldEqual, "3 Awst 2019 1:00am")  // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-08-15T00:00:00.000Z", "cy"), ShouldEqual, "15 Awst 2019 1:00am") // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-05-21T23:00:00.000Z", "cy"), ShouldEqual, "22 Mai 2019 12:00am") // BST
		So(helper.DateTimeOnsDatePatternFormat("2019-12-21T23:00:00.000Z", "cy"), ShouldEqual, "21 Rhagfyr 2019 11:00pm")
	})
}

func TestDateTimeFormat(t *testing.T) {
	Convey("Given a formatted datetime return a human readable datetime", t, func() {
		Convey("When in British Summer Time", func() {
			want := "13 June 2017 09:30"
			got := helper.DateTimeFormat("2017-06-13T08:30:00.000Z")
			So(got, ShouldEqual, want)
		})
		Convey("When not in British Summer Time (GMT)", func() {
			want := "13 February 2019 19:21"
			got := helper.DateTimeFormat("2019-02-13T19:21:22.134Z")
			So(got, ShouldEqual, want)
		})
	})
	Convey("Given a invalid datetime return said datetime", t, func() {
		want := "2006-01-02Tkjklj+07:00"
		got := helper.DateTimeFormat("2006-01-02Tkjklj+07:00")
		So(got, ShouldEqual, want)
	})
}
