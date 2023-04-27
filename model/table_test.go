package model_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTable(t *testing.T) {
	Convey("FuncContainsVariant", t, func() {
		Convey("Should detect a variant among a list of one", func() {
			table := model.Table{
				Variants: []string{"apple"},
			}
			So(table.FuncContainsVariant("apple"), ShouldBeTrue)
		})

		Convey("Should detect a variant among a list of many", func() {
			table := model.Table{
				Variants: []string{"apple", "banana", "cherry"},
			}
			So(table.FuncContainsVariant("banana"), ShouldBeTrue)
		})

		Convey("Should detect a variant missing from an empty list", func() {
			table := model.Table{
				Variants: []string{},
			}
			So(table.FuncContainsVariant("apple"), ShouldBeFalse)
		})

		Convey("Should detect a variant missing from a list of one", func() {
			table := model.Table{
				Variants: []string{"apple"},
			}
			So(table.FuncContainsVariant("apricot"), ShouldBeFalse)
		})

		Convey("Should detect a variant missing from a list of many", func() {
			table := model.Table{
				Variants: []string{"apple", "banana", "cherry"},
			}
			So(table.FuncContainsVariant("blueberry"), ShouldBeFalse)
		})
	})
}
