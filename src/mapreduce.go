package main
import (
    "fmt"
    "os"
    "mapreduce/cli"
    "mapreduce/io"
    "mapreduce/worker"
    "mapreduce/partition"
    "mapreduce/log"
)

func main() {
    fmt.Println("hello world")
    for _, val := range os.Args {
        fmt.Println(val)
    }
    var options cli.CLIOptions = cli.ParseCommandLineInput()
    fmt.Println(options.MultiProcess)
    if (options.MultiProcess) {
        fmt.Println("Multiprocess mode enabled")
    }

    success := io.FileBasedOutputStream{}.OutputData([]byte("Hello world to file! From interface implementation"), "newFile.txt")
    if success {
        fmt.Println("Successfully wrote to file")
    } else {
        fmt.Println("Failed to write to file")
    }
    data := io.FileBasedInputStream{}.RetrieveInput("inputTest.txt")
    fmt.Println(string(data))
    nilData := io.FileBasedInputStream{}.RetrieveInput("fakeFile.txt")
    fmt.Println("Nil data is: ", nilData)

    var mapWorker worker.Worker = worker.MapWorker{ TestField : "mapper"}
    var reduceWorker worker.Worker = worker.ReduceWorker{ TestField: "reducer" }

    mapWorker.Execute()
    reduceWorker.Execute()

    standardCoordinator := worker.StandardWorkerCoordinator{}
    standardCoordinator.RegisterWorker(mapWorker)
    standardCoordinator.RegisterWorker(reduceWorker)
    standardCoordinator.PrintLists()

    partitioner := partition.SequentialDataPartitioner{}
    success = partitioner.Execute([]byte("testWithStringPartitioner"))
    fetchData := partitioner.YieldData()
    fmt.Printf("Partitioned data is: %s and the identifier is: %s\n", fetchData.Data, fetchData.Identifier)

    os.Setenv("MAPREDUCE_LOG_DEBUG_ENABLED", "enabled")
    logger := log.InitializeLog(log.LogOptions{DebugEnabled: options.LogDebugEnabled})
    logger.Debug("This should only print if debug is enabled!")
    logger.Info("This should always print!")
    logger.Debug("testing out the new sequential byte array")
    var partialInputReader io.InputStream = io.NewPartialFileReader(5, 'n')
    partialData, _ := partialInputReader.RetrieveInput("lineBasedTestInput.txt")
    fmt.Println(string(partialData))
    partialData, _ = partialInputReader.RetrieveInput("lineBasedTestInput.txt")
    fmt.Println(string(partialData))

    freshDataPartitioner := partition.NewSequentialDataPartitioner(20, "\n")
    freshDataPartitioner.PartitionInput("lineBasedTestInput.txt")
    firstData := freshDataPartitioner.YieldData()
    fmt.Println("First data from partition is: ", string(firstData.Data))

    partitionedFiles := freshDataPartitioner.RetrieveInputFiles()

    for _, val := range partitionedFiles {
        standardCoordinator.RegisterInputFile(val, worker.Mapper)
    }

}
