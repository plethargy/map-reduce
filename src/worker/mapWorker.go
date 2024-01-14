package worker 
import "fmt"
type MapWorker struct {
    TestField string //this exists temporarily just to prove it works
}

func (w MapWorker) Execute() {
    fmt.Println(w.TestField)
}

func (w MapWorker) GetWorkerType() WorkerType {
    return Mapper
}
