package reduce

import (
    "fmt"
    "mapreduce/emit"
)

type ReduceInput struct {
    emitted map[string]string
}
//This is the interface that the client written reducer should adhere to.
type ReduceFunction interface {
    Reduce(r ReduceInput, e emit.Emitter)
}

type NoOpReducer struct {
}

func (nrd *NoOpReducer) Reduce(r ReduceInput, e emit.Emitter)  {
    fmt.Println("Calling no op reducer")
    return
}

type CountReducer struct {
}

func (cr *CountReducer) Reduce(r ReduceInput, e emit.Emitter) {
    //do nothing just output the count

    for k, v := range r.emitted {
        e.Emit(k, v)
    }
}

func NewReduceInput(inputData map[string]string) ReduceInput {
    return ReduceInput{emitted: inputData}
}
