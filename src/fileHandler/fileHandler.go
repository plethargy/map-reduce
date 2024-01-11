package fileHandler
//This should print the output from reduce and handle input for map. 
//For now this will just export an output function that can be called from the reducer workers.
//Lets keep it super simple (single-thread, single-execution) for now, we'll need to add consistency and rollbacks for when this is multi-process and multi-thread.
import (
    "os"
    "fmt"
)

type OutputStream interface {
    OutputData(output string, fileName string) bool
}

type FileBasedOutputStream struct { //TODO: Extract to separate files before this gets bloated
}

func (f FileBasedOutputStream) OutputData(output string, fileName string) bool {
    if !checkFileExistence(fileName) {
        createFile(fileName)
    }
    return writeToFile([]byte(output), fileName)
}


func checkFileExistence(fileName string) bool {
    _, err := os.Stat(fileName)
    if os.IsNotExist(err) {
        fmt.Println("File does not exist")
        return false
    } else {
        fmt.Println("File does not exist")
        return true
    }
}

func createFile(fileName string) {
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Println(err)
        return //at some point I should wrap this in a retryable call pattern, but that's for future Ethan :D 
    }
    defer file.Close()

    fmt.Println("Created file")
}
func writeToFile(output []byte, fileName string) bool {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening the file", err)
        return false
    }
    defer file.Close()

    _, err = file.Write(output)
    if err != nil {
        fmt.Println("Error writing output to file", err)
        return false
    }
    return true
}

