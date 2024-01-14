package io

type FileBasedOutputStream struct { //TODO: Extract to separate files before this gets bloated
}

func (f FileBasedOutputStream) OutputData(output string, fileName string) bool {
    if !checkFileExistence(fileName) {
        createFile(fileName)
    }
    return writeToFile([]byte(output), fileName)
}



