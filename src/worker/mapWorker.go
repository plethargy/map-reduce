package worker 
import "fmt"
type MapWorker struct {
    TestField string //this exists temporarily just to prove it works
}

func (w MapWorker) Execute(fileName string) {
    fmt.Println(w.TestField)
    //I envision that this will take the file input and transform it into the MapInput object that will get passed to the Map implementation.
}

func (w MapWorker) GetWorkerType() WorkerType {
    return Mapper
}
