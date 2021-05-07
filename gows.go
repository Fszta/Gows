package main

import "fmt"

func main() {
	fmt.Println("Hello gows")
}

type Operator interface {
	Run()
}

type SparkOperator struct {
	masterNodeIp string
}

func (s *SparkOperator) Run() {
	fmt.Println("spark submit")
}

func NewSparkOperator() SparkOperator {
	return SparkOperator{}
}

func NewBashOperator() *BashOperator {
	return &BashOperator{}
}

type BashOperator struct {
	scriptPath string
}

func (b *BashOperator) Run() {
	fmt.Println("launch bash")
}

func OperatorFactory(operatorName string) (Operator, error) {
	switch operatorName {
	case "bash":
		return NewBashOperator(), nil
	default:
		return nil, fmt.Errorf("BAD OPERATOR NAME")
	}
}
