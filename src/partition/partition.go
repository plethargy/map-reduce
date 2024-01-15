package partition

type PartitionedData struct {
    Data []byte
    Identifier string //if data is sequential we can use this to sort input/output
}

type DataPartitioner interface {
    Execute(dps DataPartitionStrategy, inputData []byte) PartitionedData
}

type DataPartitionStrategy interface {
    PartitionData(inputData []byte) []byte
}

type NoOpDataPartitionStrategy struct{
}

func (b NoOpDataPartitionStrategy) PartitionData(inputData []byte) []byte {
    return inputData
}

type NoOpDataPartitioner struct{}

func (dp NoOpDataPartitioner) Execute(dps DataPartitionStrategy, inputData []byte) PartitionedData {
    return PartitionedData{Data: dps.PartitionData(inputData), Identifier: "0"}
}
