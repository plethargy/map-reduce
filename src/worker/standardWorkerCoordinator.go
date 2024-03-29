package worker
import (
    "fmt"
    "sync"
    "mapreduce/map"
    "mapreduce/reduce"
)



type StandardWorkerCoordinator struct {
    MapWorkerList []Worker
    ReduceWorkerList []Worker
    MapperInputFiles []string
    ReducerInputFiles []string
    mapperFunc mapper.Mapper
    reducerFunc reduce.Reducer
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
            fmt.Println("Appended a mapper")
        }
    }
}

func (swc *StandardWorkerCoordinator) RegisterInputFile(filePath string, workerType WorkerType) {
    switch workerType {
        case Reducer: {
            swc.ReducerInputFiles = append(swc.ReducerInputFiles, filePath)
            fmt.Println("Appended a reducer file")
        }
        case Mapper: {
            swc.MapperInputFiles = append(swc.MapperInputFiles, filePath)
            fmt.Println("Appended a mapper file")
        }
    }
}

func (swc *StandardWorkerCoordinator) MapReduce(m MapReduceInput) bool {
    var wg sync.WaitGroup
    for _, worker := range swc.MapWorkerList {
        wg.Add(1)
        go func(w Worker) {
            defer wg.Done()
            w.Execute("fakefile.txt")
            swc.mapperFunc.Map(mapper.MapInput{})
        }(worker)
    }

    wg.Wait()
    return true
}

func (swc StandardWorkerCoordinator) PrintLists() {
    fmt.Println("The size of reducer list is: ", len(swc.ReduceWorkerList))
    fmt.Println("The size of mapper list is: ", len(swc.MapWorkerList))
}

func NewStandardWorkerCoordinator() StandardWorkerCoordinator {
    return StandardWorkerCoordinator{mapperFunc: &mapper.NoOpMapper{}, reducerFunc: &reduce.NoOpReducer{}}
}

func (swc *StandardWorkerCoordinator) RegisterMapper(m mapper.Mapper) {
    swc.mapperFunc = m
}

func (swc *StandardWorkerCoordinator) RegisterReducer(r reduce.Reducer) {
    swc.reducerFunc = r
}

