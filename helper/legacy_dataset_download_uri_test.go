package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLegacyDatasetDownloadURI(t *testing.T) {
	Convey("should generated expected legacy dataset download URI", t, func() {
		So(helper.LegacyDataSetDownloadURI("/legacy/dataset/page", "test.csv"), ShouldEqual, "/file?uri=/legacy/dataset/page/test.csv")
	})
}
