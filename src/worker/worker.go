package worker

import "fmt"

type WorkerType int

const (
    Reducer WorkerType = iota
    Mapper
)
type Worker interface {
    Execute()
    GetWorkerType() WorkerType
}

type Coordinator interface {
    RegisterWorker(w Worker)
    MapReduce() bool
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

func (w MapWorker) GetWorkerType() WorkerType {
    return Mapper
}
func (w ReduceWorker) Execute() {
    fmt.Println(w.TestField)
}

func (w ReduceWorker) GetWorkerType() WorkerType {
    return Reducer
}

type StandardWorkerCoordinator struct {
    MapWorkerList []Worker
    ReduceWorkerList []Worker
}
//TODO: Add a bunch more validation and logic to the appending
func (swc *StandardWorkerCoordinator) RegisterWorker(w Worker) {
    switch w.GetWorkerType() {
        case Reducer: {
            swc.ReduceWorkerList = append(swc.ReduceWorkerList, w)
            fmt.Println("Appended a reducer")
        }
        case Mapper: {
            swc.MapWorkerList = append(swc.MapWorkerList, w)
            fmt.Println("Appended a reducer")
        }
    }
    fmt.Println("The size of reducer list is: ", len(swc.ReduceWorkerList))
    fmt.Println("The size of mapper list is: ", len(swc.MapWorkerList))

}

func (swc StandardWorkerCoordinator) PrintLists() {
    fmt.Println("The size of reducer list is: ", len(swc.ReduceWorkerList))
    fmt.Println("The size of mapper list is: ", len(swc.MapWorkerList))
}
