package apibase

import (
	"github.com/kataras/iris"
	"regexp"
)

var uuidPattern = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")

func RegisterUUIDStringMacro(app *iris.Application) {
	app.Macros().String.RegisterFunc("uuid", func() func(string) bool {
		return func(pv string) bool {
			return uuidPattern.MatchString(pv)
		}
	})
}
