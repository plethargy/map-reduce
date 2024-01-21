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
    Execute(inputData []byte) bool
    YieldData() PartitionedData
    PartitionInput(fileName string) bool
}

type NoOpDataPartitionStrategy struct{
}

type NoOpDataPartitioner struct{}

func NewSequentialDataPartitioner(partitionSize int, delimter string) SequentialDataPartitioner {
    return SequentialDataPartitioner{cursor: 0, maxChunks: partitionSize, reader: io.NewPartialFileReader(10, byte('\n')) } //TODO: fix mismatch potential between partitionSize and maxChunks
}

func NewPartitionedData(data []byte, identifier string) PartitionedData {
    return PartitionedData{Data: data, Identifier: identifier}
}

func (dp NoOpDataPartitioner) Execute(inputData []byte) bool {
    return true
}

func (dp NoOpDataPartitioner) YieldData() PartitionedData {
    return PartitionedData{}
}

type SequentialDataPartitioner struct {
    ChunkedData []PartitionedData
    cursor int //when this becomes multithreaded I will need to investigate how to handle the cursor increment in an atomic way
    maxChunks int
    reader io.InputStream
}
//This should eventually take in some options input that specifies chunk size
//TODO: clean this up and add maaaaaaany error checks and modes
//TODO: This needs to be rethought. We don't want to hold all the data in memory and then chunk it, we want to rely on reading from file, this and the file reader will need to be refactored, for now lets keep this simple.
func (sdp *SequentialDataPartitioner) Execute(inputData []byte) bool {
    sdp.cursor = 0
    sdp.ChunkedData = nil
    inputLength := len(inputData)
    chunkSize := 128
    numChunks := 1 + ((inputLength - 1) / chunkSize)
    sdp.maxChunks = numChunks
    fmt.Println("Input length: ", inputLength)
    fmt.Println("Num chunks: ", chunkSize)
    for count := 0; count < numChunks; count++ {
        startIndice := chunkSize * count
        endIndice := chunkSize * (count + 1)
        if endIndice > inputLength {
            endIndice = inputLength
        }
        partitionedData := PartitionedData{Data: inputData[startIndice:endIndice], Identifier: fmt.Sprintf("%d", count)}
        sdp.ChunkedData = append(sdp.ChunkedData, partitionedData)
        fmt.Printf("Chunking data from startIndice %d to endIndice %d", startIndice, endIndice)
    }

    return true
}

func (sdp *SequentialDataPartitioner) YieldData() PartitionedData {
    if len(sdp.ChunkedData) <= 0 {
        return PartitionedData{}
    }
    if sdp.maxChunks == 0 {
        return PartitionedData{}
    }
    if sdp.cursor >= sdp.maxChunks {
        return PartitionedData{}
    }
    data := sdp.ChunkedData[sdp.cursor]
    sdp.cursor = sdp.cursor + 1
    return data
}

func (sdp *SequentialDataPartitioner) PartitionInput(fileName string) bool {
    data, err := sdp.reader.RetrieveInput(fileName)
    count := 0
    if err != nil {
        fmt.Println("Retrieved all data in first call: ", string(data))
        sdp.ChunkedData = append(sdp.ChunkedData, NewPartitionedData(data, fmt.Sprint(count))) //TODO: remove this, we don't want to store the chunked data in our partitioner, this is just for initial testing
        count++
    }
    var fetchedData []byte = nil
    for err == nil {
        if len(data) >= sdp.maxChunks {
            fmt.Println("Data received thus far: ", string(data)) //TODO: change this to write data to a smaller file
            sdp.ChunkedData = append(sdp.ChunkedData, NewPartitionedData(data, fmt.Sprint(count))) //TODO: remove this, we don't want to store the chunked data in our partitioner, this is just for initial testing
            data = nil
            count++
        }
        fetchedData, err = sdp.reader.RetrieveInput(fileName)
        data = append(data, fetchedData...)
    }
    return true
}
