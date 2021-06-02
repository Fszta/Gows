package operators

type Operator interface {
	RunTask() (string, error)
	makeCmd()
}
