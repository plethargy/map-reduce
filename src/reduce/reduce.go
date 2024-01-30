package reduce

import "fmt"

type ReduceInput struct {
}
//This is the interface that the client written reducer should adhere to.
type Reducer interface {
    Reduce(r ReduceInput)
}

type NoOpReducer struct {
}

func (nrd *NoOpReducer) Reduce(r ReduceInput) {
    fmt.Println("Calling no op reducer")
}
