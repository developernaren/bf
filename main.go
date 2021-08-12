package main

import (
	"bufio"
	"fmt"
	"naren.com/bf/pkg/bf"
	"os"
	"strings"
)

func main() {
	initialValue := "1"
	operator := bf.NewOperatorWithDefaultOperations(initialValue)

	fmt.Println("Enter the operator to operate")
	fmt.Println(fmt.Sprintf("supported operators: %s", operator.GetOperators()))
	ch := make(chan string)
	go func(ch chan string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			s, err := reader.ReadString('\n')
			if err != nil {
				close(ch)
				return
			}
			ch <- s
		}
	}(ch)

	for {
		select {
		case stdin, _ := <-ch:
			stdin = strings.TrimSpace(stdin)
			err := operator.Operate(stdin)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		fmt.Println(operator.GetValue())
	}
}
