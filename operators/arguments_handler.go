package operators

import (
	"strings"
)

type Argument struct {
	arg   string
	value string
}

type ArgumentsHandler struct {
	arguments []Argument
}

func (a *ArgumentsHandler) AddArgument(arg string, value string) {
	a.arguments = append(a.arguments, Argument{arg: arg, value: value})
}

func (a *ArgumentsHandler) getArgsToString() string {
	var argsString strings.Builder
	for _, arg := range a.arguments {
		if arg.arg != "" {
			argsString.WriteString(" " + arg.arg)
		}
		if arg.value != "" {
			argsString.WriteString(" " + arg.value)
		}
	}
	// Return the trimed stringified arguments
	return strings.TrimSpace(argsString.String())
}
