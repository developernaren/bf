## Brainfuck Task

This is a very small implementation of [brainfuck](https://en.wikipedia.org/wiki/Brainfuck)

### Implementation Details
I have written this program to only support "+" and "-". And added the functionality to add the operations as required.

### How it works

You can run the program by running

```bash
go build .
./bf
```
It will run a terminal waiting for the response

```bash
Enter the operator to operate
supported operators: +, -
```

You can enter the operator in the terminal and will operator based on the provided input


### Extendability

You can add custom operations by calling the function
The first parameter is the initial value of for the operations to be done

You can add any operation you like, it has to adhere to type `Operation`
```go
type Operation func(initialValue string) (string, error)
```

For example if you want to add an operation `^` to sqaure the value, you can just add a operation to map

```go
operations := map[string]Operation{
	"?" : func(initialValue string) (string, error) {
            intVal, err := strconv.ParseInt(value, 10, 32)
            if err != nil {
                return "", err
            }
            intVal *= intVal;
            return fmt.Sprintf("%d", intVal), nil	
    }   
}
NewOperator(initivalValue, operations)
```

#### Why not `addOperation` in `Operator`

Because I want our Operator to be immutable

### How to run tests

Run `go test -count=1 ./...` to run tests
