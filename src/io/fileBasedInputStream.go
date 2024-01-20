package io
import ( 
    "fmt"
    "os"
    "bufio"
    "io"
)
var (
    newlineByteRepresentation byte = byte('\n')
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


type PartialReadFileBasedInputStream struct {
    LinesToFetch int
    fileOpen bool
    bufferedReader *bufio.Reader
}

func (p *PartialReadFileBasedInputStream) RetrieveInput(fileName string) []byte {
    if !checkFileExistence(fileName) {
        return nil
    }
    var file *os.File
    if !p.fileOpen {
        file, _ = os.Open(fileName)
        p.bufferedReader = bufio.NewReader(file)
        p.fileOpen = true
    }
    byteArray := make([]byte, 0)
    for i := 0; i < p.LinesToFetch; i++ {
        readData, err := p.bufferedReader.ReadBytes('\n')
        byteArray = append(byteArray, readData...)
        if err != nil {
            if err == io.EOF {
                defer file.Close()
                p.fileOpen = false
            } else {
                fmt.Println("Something fatal occurred reading in input", err)
            }
        }
    }
    return byteArray
}

func NewPartialFileReader(linesToFetchPerExecution int) PartialReadFileBasedInputStream {
    return PartialReadFileBasedInputStream{LinesToFetch: linesToFetchPerExecution, fileOpen: false}
}
