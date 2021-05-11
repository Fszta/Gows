package operators

import (
	"encoding/json"
	"fmt"
)

type IOperator interface {
	RunTask()
}

type Operator struct {
	Name       string          `json:"name"`
	Parameters json.RawMessage `json:"parameters"`
}

func NewOperator(operatorName string) (IOperator, error) {
	switch operatorName {
	case "bash":
		return NewBashOperator(), nil
	default:
		return nil, fmt.Errorf("OPERATOR %s DOESN'T EXISTS", operatorName)
	}
}
