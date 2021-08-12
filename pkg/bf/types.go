package bf

type Operation func(initialValue string) (string, error)

type Operator struct {
	operations map[string]Operation
	value string
}
