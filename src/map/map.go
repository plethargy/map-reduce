package mapper

import (
	"fmt"
)

type MapInput struct {
    data []byte
}
//This is the interface intended to shape the function that client written Map functions should adhere to.
type MapFunction interface {
    Map(m MapInput) string
}

type NoOpMapper struct {
}

func (nmp *NoOpMapper) Map(m MapInput) string {
    fmt.Println("Calling No Op Mapper")
    var defaultResult string
    return defaultResult
}

func NewMapInput(inputData []byte) MapInput {
    return MapInput{data: inputData}
}
