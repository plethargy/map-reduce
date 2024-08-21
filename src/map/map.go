package mapper

import (
	"fmt"
	"mapreduce/emit"
	"strconv"
	"strings"
)

type MapInput struct {
    data []byte
}
//This is the interface intended to shape the function that client written Map functions should adhere to.
type MapFunction interface {
    Map(m MapInput, e emit.Emitter)
}

type NoOpMapper struct {
}

func (nmp *NoOpMapper) Map(m MapInput, e emit.Emitter) {
    fmt.Println("Calling No Op Mapper")
    return 
}

func NewMapInput(inputData []byte) MapInput {
    return MapInput{data: inputData}
}

type CountMapper struct {
}

func (cm  *CountMapper) Map(m MapInput, e emit.Emitter) {
    lines := strings.Split(string(m.data), "\n")
    countMap := make(map[string]int)
    for _, elem := range lines {
	words := strings.Split(elem, " ")
	for _, word := range words {
	    value, exists := countMap[word]
	    if exists {
		countMap[word] = value + 1
	    } else {
		countMap[word] = 1
	    }
	}
    }
    for k, v := range countMap {
	e.Emit(k, strconv.Itoa(v))
    }
}
