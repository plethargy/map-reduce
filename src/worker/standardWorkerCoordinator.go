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
    mapperFunc mapper.MapFunction
    reducerFunc reduce.ReduceFunction
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
    // we want to ensure every file gets handled before moving to the Reducers 
    // this will iterate through each file and hand it to a MapWorker if one is Idle, if they're all busy it will spin
    // until one becomes Idle.
    var fileIndex = 0
    var wg sync.WaitGroup
    inputFileLength := len(swc.MapperInputFiles)
    for fileIndex < inputFileLength {
        for _, worker := range swc.MapWorkerList {
            if worker.GetWorkerStatus() == Idle && fileIndex < inputFileLength {
                wg.Add(1)
                worker.SetWorkerStatus(Busy)
                go func(w Worker, indx int) {
                    defer wg.Done()
                    worker.ExecuteMap(swc.MapperInputFiles[indx], swc.mapperFunc)
                    //swc.mapperFunc.Map(mapper.MapInput{})
                }(worker, fileIndex)
                fileIndex += 1
            }
        }
        wg.Wait()
    }
    return true
}

func (swc StandardWorkerCoordinator) PrintLists() {
    fmt.Println("The size of reducer list is: ", len(swc.ReduceWorkerList))
    fmt.Println("The size of mapper list is: ", len(swc.MapWorkerList))
}

func NewStandardWorkerCoordinator() StandardWorkerCoordinator {
    return StandardWorkerCoordinator{mapperFunc: &mapper.NoOpMapper{}, reducerFunc: &reduce.NoOpReducer{}}
}

func (swc *StandardWorkerCoordinator) RegisterMapper(m mapper.MapFunction) {
    swc.mapperFunc = m
}

func (swc *StandardWorkerCoordinator) RegisterReducer(r reduce.ReduceFunction) {
    swc.reducerFunc = r
}

