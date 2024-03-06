package model_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFuncHasDateValidationErr(t *testing.T) {
	Convey("Given a date input fieldset", t, func() {
		Convey("When the FuncHasDateValidationErr function is called", func() {
			Convey("Then it returns the expected bool", func() {
				tc := []struct {
					given    model.DateFieldset
					expected bool
				}{
					{
						given:    model.DateFieldset{},
						expected: false,
					},
					{
						given: model.DateFieldset{
							Input: model.InputDate{
								HasDayValidationErr: true,
							},
						},
						expected: true,
					},
					{
						given: model.DateFieldset{
							Input: model.InputDate{
								HasMonthValidationErr: true,
							},
						},
						expected: true,
					},
					{
						given: model.DateFieldset{
							Input: model.InputDate{
								HasYearValidationErr: true,
							},
						},
						expected: true,
					},
					{
						given: model.DateFieldset{
							Input: model.InputDate{
								HasYearValidationErr:  true,
								HasMonthValidationErr: false,
								HasDayValidationErr:   true,
							},
						},
						expected: true,
					},
				}
				for _, t := range tc {
					So(t.given.FuncHasDateValidationErr(), ShouldEqual, t.expected)
				}
			})
		})
	})
}
