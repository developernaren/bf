package bf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOperatorWithDefaultOperations(t *testing.T) {

	operator := NewOperatorWithDefaultOperations("1")

	err := operator.Operate("+")
	assert.Nil(t, err)
	assert.Equal(t, "2", operator.GetValue())
	assert.Contains(t,  operator.GetOperators(), "+",)
	assert.Contains(t, operator.GetOperators(), "-")
}

func TestNewOperator(t *testing.T) {

	operations := map[string]Operation{
		"?" : func(initialValue string) (string, error) {
			return "10", nil
		},
	}
	operator := NewOperator("1", operations )
	err := operator.Operate("?")

	assert.Nil(t, err)
	assert.Equal(t, "10", operator.GetValue())
	assert.Equal(t, "?", operator.GetOperators())
}

func TestUndefinedOperationThrowsError(t *testing.T) {

	operator := NewOperatorWithDefaultOperations("1")

	err := operator.Operate("?")

	assert.Error(t, err)
	assert.Equal(t, "operator '?' is not supported", err.Error())
}
