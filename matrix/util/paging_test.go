package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPaging(t *testing.T) {
	Convey("Test Paging", t, func() {
		data := []interface{}{"a", "b", "c", "d"}
		So(Paginate(data, 1, 2), ShouldResemble, []interface{}{"a", "b"})
		So(Paginate(data, 2, 2), ShouldResemble, []interface{}{"c", "d"})
		So(Paginate(data, 1, 3), ShouldResemble, []interface{}{"a", "b", "c"})
		So(Paginate(data, 2, 3), ShouldResemble, []interface{}{"d"})
	})
}
