package account

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateTelephone(t *testing.T) {
	Convey("test telephone", t, func() {
		So(ValidateTelephone("13177905432"), ShouldBeTrue)
		So(ValidateTelephone("15812344321"), ShouldBeTrue)
		So(ValidateTelephone("18977905432"), ShouldBeTrue)
		So(ValidateTelephone("131-77905432"), ShouldBeFalse)
		So(ValidateTelephone("+13177905432"), ShouldBeFalse)
	})
}

func TestValidateEmail(t *testing.T) {
	Convey("test email", t, func() {
		So(ValidateEmail("hatlonely@foxmail.com"), ShouldBeTrue)
		So(ValidateEmail("abc.xyz@abc-def.com"), ShouldBeTrue)
		So(ValidateEmail("abc$def.com"), ShouldBeFalse)
		So(ValidateEmail("ç$€§/az@gmail.com"), ShouldBeFalse)
	})
}
