package io

type FileBasedOutputStream struct { //TODO: Extract to separate files before this gets bloated
}

func (f FileBasedOutputStream) OutputData(output []byte, fileName string) bool {
    if !checkFileExistence(fileName) {
        createFile(fileName)
    }
    return writeToFile(output, fileName)
}



