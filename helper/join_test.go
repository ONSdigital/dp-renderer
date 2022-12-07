package helper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJoin(t *testing.T) {
	Convey("Join returns a concatenated string from a slice", t, func() {
		input := []string{"alpha", "beta", "gamma"}

		got := Join(",", input)
		want := "alpha,beta,gamma"
		So(got, ShouldEqual, want)
	})
}
