package worker
import (
    "fmt"
    "sync"
    "mapreduce/map"
    "mapreduce/reduce"
)



type StandardWorkerCoordinator[T any] struct {
    MapWorkerList []Worker
    ReduceWorkerList []Worker
    MapperInputFiles []string
    ReducerInputFiles []string
    mapperFunc mapper.Mapper[T]
    reducerFunc reduce.Reducer[T]
}
//TODO: Add a bunch more validation and logic to the appending
func (swc *StandardWorkerCoordinator[T]) RegisterWorker(w Worker) {
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

func (swc *StandardWorkerCoordinator[T]) RegisterInputFile(filePath string, workerType WorkerType) {
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

func (swc *StandardWorkerCoordinator[T]) MapReduce(m MapReduceInput) bool {
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

func (swc StandardWorkerCoordinator[T]) PrintLists() {
    fmt.Println("The size of reducer list is: ", len(swc.ReduceWorkerList))
    fmt.Println("The size of mapper list is: ", len(swc.MapWorkerList))
}

func NewStandardWorkerCoordinator() StandardWorkerCoordinator[string] {
    return StandardWorkerCoordinator[string]{mapperFunc: &mapper.NoOpMapper[string]{}, reducerFunc: &reduce.NoOpReducer[string]{}}
}

func (swc *StandardWorkerCoordinator[T]) RegisterMapper(m mapper.Mapper[T]) {
    swc.mapperFunc = m
}

func (swc *StandardWorkerCoordinator[T]) RegisterReducer(r reduce.Reducer[T]) {
    swc.reducerFunc = r
}

