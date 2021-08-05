package greeting

import (
	"strings"
)

func Greet(name string) string {
	if name == strings.ToUpper(name) {
		return "Capital"
	}
	return "Lower"
}
