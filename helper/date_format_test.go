package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDateFormat(t *testing.T) {
	Convey("Date format returns human readable string", t, func() {
		So(helper.DateFormat("2019-08-15T00:00:00.000Z"), ShouldEqual, "15 August 2019")
		So(helper.DateFormat("2019-08-15"), ShouldEqual, "2019-08-15")
		So(helper.DateFormat(""), ShouldEqual, "")
	})
}

func TestDateFormatYYYYMMDD(t *testing.T) {
	Convey("Date format returns short date pattern without slashes", t, func() {
		So(helper.DateFormatYYYYMMDD("2019-08-15T00:00:00.000Z"), ShouldEqual, "2019/08/15")
		So(helper.DateFormatYYYYMMDD("2019-08-15"), ShouldEqual, "2019-08-15")
		So(helper.DateFormatYYYYMMDD(""), ShouldEqual, "")
	})
}

func TestDateFormatYYYYMMDDNoSlash(t *testing.T) {
	Convey("Date format returns human readable string", t, func() {
		So(helper.DateFormatYYYYMMDDNoSlash("2019-08-15T00:00:00.000Z"), ShouldEqual, "20190815")
		So(helper.DateFormatYYYYMMDDNoSlash("2019-08-15"), ShouldEqual, "2019-08-15") // failed to parse, so returns arg value
		So(helper.DateFormatYYYYMMDDNoSlash(""), ShouldEqual, "")
	})
}

func TestDateTimeFormat(t *testing.T) {
	Convey("Given a formatted datetime return a human readable datetime", t, func() {
		want := "13 June 2017 08:30"
		got := helper.DateTimeFormat("2017-06-13T08:30:00.000Z")
		So(got, ShouldEqual, want)
	})
	Convey("Given a invalid datetime return said datetime", t, func() {
		want := "2006-01-02Tkjklj+07:00"
		got := helper.DateTimeFormat("2006-01-02Tkjklj+07:00")
		So(got, ShouldEqual, want)
	})
}
