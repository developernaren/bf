package bf

import (
	"fmt"
	"strconv"
)

var addOperation = func(value string) (string, error){
	intVal, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return "", err
	}
	intVal += 1;
	return fmt.Sprintf("%d", intVal), nil
}

var subOperation = func(value string) (string, error) {
	intVal, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return "", err
	}
	intVal -= 1;
	return fmt.Sprintf("%d", intVal), nil
}

func GetDefaultOperations() map[string]Operation {

	return map[string]Operation{
		"+" : addOperation,
		"-" : subOperation,
	}
}
