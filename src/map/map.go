package mapper

import "fmt"

type MapInput struct {
}
//This is the interface intended to shape the function that client written Map functions should adhere to.
type Mapper interface {
    Map(m MapInput)
}
//TODO: Figure out how to get the Emit to emit to the coordinator

type NoOpMapper struct {
}

func (nmp *NoOpMapper) Map(m MapInput) {
    fmt.Println("Calling No Op Mapper")
}
