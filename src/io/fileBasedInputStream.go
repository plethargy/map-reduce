package io

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)
var (
    newlineByteRepresentation byte = byte('\n')
)
type FileBasedInputStream struct {
}
func (f FileBasedInputStream) RetrieveInput(fileName string) ([]byte, error) {
    if !checkFileExistence(fileName) {
        return nil, errors.New("File does not exist")
    }
    data, err := os.ReadFile(fileName) //this will definitely need changing to support larger volumes of data, reading it all into memory at once is bad
    if err != nil {
        fmt.Println("Error retrieving data", err)
        return nil, err
    }
    return data, nil
}


type PartialReadFileBasedInputStream struct {
    LinesToFetch int
    fileOpen bool
    bufferedReader *bufio.Reader
    delimiter byte
}

func (p *PartialReadFileBasedInputStream) RetrieveInput(fileName string) ([]byte, error) {
    if !checkFileExistence(fileName) {
        return nil, errors.New("File does not exist")
    }
    var file *os.File
    if !p.fileOpen {
        file, _ = os.Open(fileName)
        p.bufferedReader = bufio.NewReader(file)
        p.fileOpen = true
    }
    byteArray := make([]byte, 0)
    for i := 0; i < p.LinesToFetch; i++ {
        readData, err := p.bufferedReader.ReadBytes(p.delimiter)
        byteArray = append(byteArray, readData...)
        if err != nil {
            if err == io.EOF {
                defer file.Close()
                p.fileOpen = false
                return byteArray, err
            } else {
                defer file.Close()
                fmt.Println("Something fatal occurred reading in input", err)
            }
        }
    }
    return byteArray, nil
}

func NewPartialFileReader(linesToFetchPerExecution int, delimiter byte) InputStream {
    return &PartialReadFileBasedInputStream{LinesToFetch: linesToFetchPerExecution, fileOpen: false, delimiter: delimiter}
}

func NewFileReader() InputStream {
    return &FileBasedInputStream{}
}
