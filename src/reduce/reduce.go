package reduce

import "fmt"

type ReduceInput struct {
}
//This is the interface that the client written reducer should adhere to.
type ReduceFunction interface {
    Reduce(r ReduceInput) string
}

type NoOpReducer struct {
}

func (nrd *NoOpReducer) Reduce(r ReduceInput) string {
    fmt.Println("Calling no op reducer")
    var defaultResult string
    return defaultResult
}
