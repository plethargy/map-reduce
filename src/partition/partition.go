package partition

import (
    "fmt"
    "mapreduce/io"
)
type PartitionedData struct {
    Data []byte
    Identifier string //if data is sequential we can use this to sort input/output
}

type DataPartitioner interface {
    PartitionInput(fileName string) bool
}

func NewSequentialDataPartitioner(partitionSize int, delimter string) SequentialDataPartitioner {
    return SequentialDataPartitioner{maxChunks: partitionSize, reader: io.NewPartialFileReader(10, byte('\n')), writer: io.FileBasedOutputStream{} } //TODO: fix mismatch potential between partitionSize and maxChunks
}

func NewPartitionedData(data []byte, identifier string) PartitionedData {
    return PartitionedData{Data: data, Identifier: identifier}
}

type SequentialDataPartitioner struct {
    maxChunks int
    reader io.InputStream
    writer io.OutputStream
    partionedFiles []string
}


func (sdp *SequentialDataPartitioner) PartitionInput(fileName string) bool {
    data, err := sdp.reader.RetrieveInput(fileName)
    count := 0
    if err != nil {
        fmt.Println("Retrieved all data in first call: ", string(data))
        sdp.outputChunkedData(data, count, "chunkedData")
    }
    var fetchedData []byte = nil
    for err == nil {
        if len(data) >= sdp.maxChunks {
            fmt.Println("Data received thus far: ", string(data))
            sdp.outputChunkedData(data, count, "chunkedData")
            data = nil
            count++
        }
        fetchedData, err = sdp.reader.RetrieveInput(fileName)
        data = append(data, fetchedData...)
    }
    return true
}

func (sdp *SequentialDataPartitioner) outputChunkedData(data []byte, count int, filePattern string) {
    countStringFormat := fmt.Sprint(count)
    fileName := fmt.Sprintf("%s%s.txt", filePattern, countStringFormat)
    sdp.writer.OutputData(data, fileName)
    sdp.partionedFiles = append(sdp.partionedFiles, fileName) 
}

func (sdp SequentialDataPartitioner) RetrieveInputFiles() []string {
    return sdp.partionedFiles
}
