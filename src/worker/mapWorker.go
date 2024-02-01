package worker 
import (
    "fmt"
)
type MapWorker struct {
    TestField string //this exists temporarily just to prove it works
}

func (w MapWorker) Execute(fileName string) {
    fmt.Println(w.TestField)
    //I envision that this will take the file input and transform it into the MapInput object that will get passed to the Map implementation.
    //TODO: get data from fileName and parse it into a MapInput object
}

func (w MapWorker) GetWorkerType() WorkerType {
    return Mapper
}
