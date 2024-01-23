package worker
import "fmt"

type StandardWorkerCoordinator struct {
    MapWorkerList []Worker
    ReduceWorkerList []Worker
    MapperInputFiles []string
    ReducerInputFiles []string
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

func (swc StandardWorkerCoordinator) PrintLists() {
    fmt.Println("The size of reducer list is: ", len(swc.ReduceWorkerList))
    fmt.Println("The size of mapper list is: ", len(swc.MapWorkerList))
}
