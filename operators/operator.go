package operators

type Operator interface {
	RunTask() (string, error)
	SetCmd(cmd string)
	SetArguments(args []Argument)
}
