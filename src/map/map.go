package mapper

import "fmt"

type MapInput struct {
}
//This is the interface intended to shape the function that client written Map functions should adhere to.
type Mapper[T any] interface {
    Map(m MapInput) T
}

type NoOpMapper[T any] struct {
}

func (nmp *NoOpMapper[T]) Map(m MapInput) T {
    fmt.Println("Calling No Op Mapper")
    var defaultResult T 
    return defaultResult
}
