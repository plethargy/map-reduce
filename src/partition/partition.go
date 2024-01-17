package partition

import "fmt"

type PartitionedData struct {
    Data []byte
    Identifier string //if data is sequential we can use this to sort input/output
}

type DataPartitioner interface {
    Execute(inputData []byte) bool
    YieldData() PartitionedData
}

type NoOpDataPartitionStrategy struct{
}

type NoOpDataPartitioner struct{}

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
}
//This should eventually take in some options input that specifies chunk size
//TODO: clean this up and add maaaaaaany error checks and modes
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
