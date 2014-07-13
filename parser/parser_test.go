package parser

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJsonParser(t *testing.T) {

	Convey("Given a received JSON", t, func() {

		Convey("When JSON is valid", func() {

			received := &Received{
				RedirectUri: "http://google.com",
				Resource:    "{}",
			}

			Convey("The JSON should be parsed", func() {
				So(Parse("{\"redirectUri\": \"http://google.com\", \"resource\": \"{}\"}"), ShouldResemble, received)
			})

		})

		Convey("When JSON is invalid", func() {

			Convey("The answer should be an error", nil)

		})

	})

}
