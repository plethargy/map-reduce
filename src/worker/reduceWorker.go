package worker

import (
    "mapreduce/reduce"
    "mapreduce/map"
    "fmt"
)

type ReduceWorker struct {

    Status WorkerStatus
    TestField string //same as above
}

func (w *ReduceWorker) ExecuteReduce(fileName string, reduceFunc reduce.ReduceFunction) {
    fmt.Println(w.TestField)
    w.Status = Busy
}

func (w ReduceWorker) GetWorkerType() WorkerType {
    return Reducer
}
func (w ReduceWorker) GetWorkerStatus() WorkerStatus {
    return w.Status
}

func (w *ReduceWorker) SetWorkerStatus(status WorkerStatus) {
    w.Status = status
}

func (w *ReduceWorker) ExecuteMap(fileName string, mapFunc mapper.MapFunction) {
    return // similar to map, we need to rethink this abstraction
}
