package worker

import "fmt"

type Worker interface {
    Execute()
}

type Coordinator interface {
    RegisterReduceWorker(w *Worker)
    MapReduce() bool
    RegisterMapWorker(w *Worker) //TODO: Revisit this, I don't like having two separate functions, but I don't want to import an entire package just to do reflection on the types.
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

type StandardWorkerCoordinator struct {
    MapWorkerList []Worker
    ReduceWorkerList []Worker
}
//TODO: Add a bunch more validation and logic to the appending
func (swc StandardWorkerCoordinator) RegisterReduceWorker(w * Worker) {
    swc.ReduceWorkerList = append(swc.ReduceWorkerList, *w)
    fmt.Println("Appended a reducer")
}

func (swc StandardWorkerCoordinator) RegisterMapWorker(w *Worker) {
    swc.MapWorkerList = append(swc.MapWorkerList, *w)
    fmt.Println("Appended a mapper")
}
