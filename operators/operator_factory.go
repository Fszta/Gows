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

func NewOperator(data []byte) (IOperator, error) {
	operator := UnmarshalOperator(data)

	switch operator.Name {
	case "bash":
		return NewBashOperator(operator.Parameters), nil
	default:
		return nil, fmt.Errorf("OPERATOR %s DOESN'T EXISTS", operator.Name)
	}
}

func UnmarshalOperator(data []byte) Operator {
	var taskConf = Operator{}

	err := json.Unmarshal(data, &taskConf)

	if err != nil {
		panic(err)
	}

	return taskConf
}
