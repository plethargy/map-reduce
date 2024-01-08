package worker

import "fmt"

type Worker interface {
    Execute()
}

type MapWorker struct {
    TestField string //this exists temporarily just to prove it works
}

type ReduceWorker struct {
    TestField string //same as above
}

func (w MapWorker) Execute() {
    fmt.Println(w.TestField)
}

func (w ReduceWorker) Execute() {
    fmt.Println(w.TestField)
}
