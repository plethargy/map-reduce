package worker

import "fmt"

type ReduceWorker struct {

    TestField string //same as above
}

func (w ReduceWorker) Execute() {
    fmt.Println(w.TestField)
}

func (w ReduceWorker) GetWorkerType() WorkerType {
    return Reducer
}
