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
    ExecuteMap(fileName string, mapFunc mapper.MapFunction) string
    ExecuteReduce(fileName string, reduceFunc reduce.ReduceFunction) string
    GetWorkerType() WorkerType
    GetWorkerStatus() WorkerStatus
    SetWorkerStatus(status WorkerStatus)
}

type Coordinator interface {
    RegisterWorker(w Worker)
    MapReduce() bool
    RegisterInputFile(fileName string, workerType WorkerType)
    RegisterMapper(m mapper.MapFunction)
    RegisterReducer(r reduce.ReduceFunction)
}
