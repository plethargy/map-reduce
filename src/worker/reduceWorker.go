package worker

import (
	"fmt"
	"mapreduce/io"
	"mapreduce/map"
	"mapreduce/reduce"
        "mapreduce/emit"
	"strings"
)

type ReduceWorker struct {
    identifier string //same as above
    Status WorkerStatus
    outputFilePath string
}

func (w *ReduceWorker) ExecuteReduce(fileName string, reduceFunc reduce.ReduceFunction) string {
    fmt.Println(w.identifier)
    w.Status = Busy
    data, err := io.NewFileReader().RetrieveInput(fileName)
    if err != nil {
        return ""
    }
    lines := strings.Split(string(data), "\n")

    inputMap := make(map[string]string)
    for _, elem := range lines {
        lineItem := strings.Split(elem, ",")
        if len(lineItem) > 1 {
            inputMap[lineItem[0]] = lineItem[1]
        }
    }
    emitter := emit.NewEmitter()
    reduceInput := reduce.NewReduceInput(inputMap)
    reduceFunc.Reduce(reduceInput, emitter)
    fileOutput := fmt.Sprintf("%s_%s.txt", w.outputFilePath, w.identifier)
    emitter.WriteDataBuffer(io.FileBasedOutputStream{}, fileOutput)

    w.Status = Idle // should probably pull this into a function and defer it
    
    return fileOutput
}

func (w ReduceWorker) GetWorkerType() WorkerType {
    return Reducer
}
func (w ReduceWorker) GetWorkerStatus() WorkerStatus {
    return w.Status
}

func (w *ReduceWorker) SetWorkerStatus(status WorkerStatus) {
    w.Status = status
}

func (w *ReduceWorker) ExecuteMap(fileName string, mapFunc mapper.MapFunction) string {
    return ""// similar to map, we need to rethink this abstraction
}

func NewReduceWorker(id string, outputFilePath string) Worker {
    return &ReduceWorker{identifier: id, Status: Idle, outputFilePath: outputFilePath}
}
