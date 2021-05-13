package operators

import (
	"fmt"
)

func NewOperator(operatorName string) (Operator, error) {
	switch operatorName {
	case "bash":
		return NewBashOperator(), nil
	default:
		return nil, fmt.Errorf("OPERATOR %s DOESN'T EXISTS", operatorName)
	}
}
