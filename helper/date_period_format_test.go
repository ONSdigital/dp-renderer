package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatePeriodFormat(t *testing.T) {
	Convey("Given a time-series monthly string", t, func() {
		want := "Jun - Sep 2010"
		got := helper.DatePeriodFormat("2010 JUN-SEP")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series monthly string spanning two different years", t, func() {
		want := "Dec 2010 - Jan 2011"
		got := helper.DatePeriodFormat("2010 DEC-JAN")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series yearly string", t, func() {
		want := "2019"
		got := helper.DatePeriodFormat("2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series month string", t, func() {
		want := "Feb 2018"
		got := helper.DatePeriodFormat("2018 FEB")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q1 string", t, func() {
		want := "Jan - Mar 2019"
		got := helper.DatePeriodFormat("Q1 2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q2 string", t, func() {
		want := "Apr - Jun 2019"
		got := helper.DatePeriodFormat("Q2 2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q3 string", t, func() {
		want := "Jul - Sep 2019"
		got := helper.DatePeriodFormat("Q3 2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q4 string", t, func() {
		want := "Oct - Dec 2019"
		got := helper.DatePeriodFormat("Q4 2019")
		So(got, ShouldEqual, want)
	})
}
