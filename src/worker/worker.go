package worker

import (
    "mapreduce/map"
    "mapreduce/reduce"
)

type WorkerType int

const (
    Reducer WorkerType = iota
    Mapper
)
type Worker interface {
    Execute(fileName string)
    GetWorkerType() WorkerType
}

type Coordinator[T any] interface {
    RegisterWorker(w Worker)
    MapReduce(m MapReduceInput) bool
    RegisterInputFile(fileName string, workerType WorkerType)
    RegisterMapper(m mapper.Mapper[T])
    RegisterReducer(r reduce.Reducer[T])
}

type MapReduceInput struct {
    intermediateFilePath string
    inputFile string
    outputFile string
}

func NewMapReduceInput(intermediateFilePath string, inputFile string, outputFile string) MapReduceInput {
    return MapReduceInput{intermediateFilePath: intermediateFilePath, inputFile: inputFile, outputFile: outputFile}
}
