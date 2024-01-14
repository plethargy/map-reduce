package io
import ( 
    "fmt"
    "os"
)
type FileBasedInputStream struct {
}

func (f FileBasedInputStream) RetrieveInput(fileName string) []byte {
    if !checkFileExistence(fileName) {
        return nil
    }
    data, err := os.ReadFile(fileName) //this will definitely need changing to support larger volumes of data, reading it all into memory at once is bad
    if err != nil {
        fmt.Println("Error retrieving data", err)
        return nil
    }
    return data
}


