package worker 
import (
    "fmt"
    "mapreduce/map"
    "mapreduce/reduce"
    "mapreduce/io"
)
type MapWorker struct {
    TestField string //this exists temporarily just to prove it works
    Status WorkerStatus

}

func (w *MapWorker) ExecuteMap(fileName string, mapFunc mapper.MapFunction) {
    fmt.Println(w.TestField)
    //I envision that this will take the file input and transform it into the MapInput object that will get passed to the Map implementation.
    //TODO: get data from fileName and parse it into a MapInput object
    w.Status = Busy
    fileReader := io.NewFileReader() // the files have already been partitioned so we can just load all the data
    mapData, err := fileReader.RetrieveInput(fileName)

    if err != nil {
        return // we should actually return a status and let the coordinator track failed attempts but that will be in a future iteration
    }

    mapFunc.Map(mapper.NewMapInput(mapData))

    w.Status = Idle

}

func (w MapWorker) GetWorkerType() WorkerType {
    return Mapper
}

func (w MapWorker) GetWorkerStatus() WorkerStatus {
    return w.Status
}

func (w *MapWorker) SetWorkerStatus(status WorkerStatus) {
    w.Status = status
}

func (w *MapWorker) ExecuteReduce(fileName string, reduceFunc reduce.ReduceFunction) {
    return // I'm not a fan of this as it technically breaks SOLID principles but I'll refactor this further down the line
}
