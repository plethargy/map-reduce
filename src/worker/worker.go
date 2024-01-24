package worker

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
    RegisterInputFile(fileName string, workerType WorkerType)
}

