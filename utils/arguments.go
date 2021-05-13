package operators

import (
	"strings"
)

type Argument struct {
	Arg   string
	Value string
}

func GetArgsToString(params []Argument) string {
	var argsString strings.Builder
	for _, arg := range params {
		if arg.Arg != "" {
			argsString.WriteString(" " + arg.Arg)
		}
		if arg.Value != "" {
			argsString.WriteString(" " + arg.Value)
		}
	}
	// Return the trimed stringified arguments
	return strings.TrimSpace(argsString.String())
}
