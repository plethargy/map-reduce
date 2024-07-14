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

type WorkerStatus int 

const(
    Idle WorkerStatus = iota
    Busy
)

type Worker interface {
    ExecuteMap(fileName string, mapFunc mapper.MapFunction)
    ExecuteReduce(fileName string, reduceFunc reduce.ReduceFunction)
    GetWorkerType() WorkerType
    GetWorkerStatus() WorkerStatus
    SetWorkerStatus(status WorkerStatus)
}

type Coordinator interface {
    RegisterWorker(w Worker)
    MapReduce(m MapReduceInput) bool
    RegisterInputFile(fileName string, workerType WorkerType)
    RegisterMapper(m mapper.MapFunction)
    RegisterReducer(r reduce.ReduceFunction)
}

type MapReduceInput struct {
    intermediateFilePath string
    inputFile string
    outputFile string
}

func NewMapReduceInput(intermediateFilePath string, inputFile string, outputFile string) MapReduceInput {
    return MapReduceInput{intermediateFilePath: intermediateFilePath, inputFile: inputFile, outputFile: outputFile}
}
