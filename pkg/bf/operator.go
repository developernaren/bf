package bf

import (
	"fmt"
	"strings"
)

func NewOperator(initialValue string, operations map[string]Operation) Operator {
	return Operator{operations, initialValue}
}

func NewOperatorWithDefaultOperations(initialValue string) Operator {
	return Operator{GetDefaultOperations(), initialValue}
}

func (operator *Operator) Operate(char string) error {

	char = strings.TrimSpace(char)
	if operation, ok := operator.operations[char]; ok {
		operatedValue, err := operation(operator.value)
		if err != nil {
			return err
		}
		operator.value = operatedValue
		return nil
	}

	return fmt.Errorf("operator '%s' is not supported", char)
}

func (operator *Operator) GetValue() string {
	return operator.value
}

func (operator *Operator) GetOperators() string {
	var operators []string

	for operator :=  range operator.operations {
		operators =  append(operators, operator)
	}

	return strings.Join(operators, ", ")
}
