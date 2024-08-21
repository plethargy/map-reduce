package emit

import (
    "mapreduce/io"
)

type Emitter struct {
    emittedData map[string]string
}

func (e* Emitter) Emit(key string, value string) {
    e.emittedData[key] = value
}

func (e* Emitter) WriteDataBuffer(outputStream io.OutputStream, fileName string) {
    dataToWrite := ""
    for k, v := range e.emittedData {
        dataToWrite += k + "," + v + "\n"
    }
    outputStream.OutputData([]byte(dataToWrite), fileName)
}

func NewEmitter() Emitter {
    return Emitter{emittedData: make(map[string]string)}
}
