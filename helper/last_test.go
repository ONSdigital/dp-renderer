package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLast(t *testing.T) {
	cases := []struct {
		Description string
		Index       int
		Data        interface{}
		Expected    bool
	}{
		{"True: 2 is the last index value of of ['a', 'b', 'c']", 2, []string{"a", "b", "c"}, true},
		{"True: 2 is the last index value of 'abc'", 2, "abc", true},
		{"True: 1 is the last index value of ['1', '2']", 1, []string{"1", "2"}, true},
		{"False: 5 is not the length of an int array with 2 elements", 5, []int{1, 2}, false},
		{"False: 3 is last index value of 'abcdefgh", 3, "abcdefgh", false},
	}

	for _, test := range cases {
		Convey(test.Description, t, func() {
			got := helper.Last(test.Index, test.Data)
			So(got, ShouldEqual, test.Expected)
		})
	}
}
