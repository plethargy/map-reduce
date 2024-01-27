package worker

type WorkerType int

const (
    Reducer WorkerType = iota
    Mapper
)
type Worker interface {
    Execute(fileName string)
    GetWorkerType() WorkerType
}

type Coordinator interface {
    RegisterWorker(w Worker)
    MapReduce(m MapReduceInput) bool
    RegisterInputFile(fileName string, workerType WorkerType)
}

type MapReduceInput struct {
    intermediateFilePath string
    inputFile string
    outputFile string
}

func NewMapReduceInput(intermediateFilePath string, inputFile string, outputFile string) MapReduceInput {
    return MapReduceInput{intermediateFilePath: intermediateFilePath, inputFile: inputFile, outputFile: outputFile}
}
