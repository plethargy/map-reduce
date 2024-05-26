package reduce

import "fmt"

type ReduceInput struct {
}
//This is the interface that the client written reducer should adhere to.
type Reducer[T any] interface {
    Reduce(r ReduceInput) T
}

type NoOpReducer[T any] struct {
}

func (nrd *NoOpReducer[T]) Reduce(r ReduceInput) T {
    fmt.Println("Calling no op reducer")
    var defaultResult T 
    return defaultResult
}
