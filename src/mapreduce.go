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
    for _, val := range os.Args {
        fmt.Println(val)
    }
    var options cli.CLIOptions = cli.ParseCommandLineInput()
    if (options.MultiProcess) {
        fmt.Println("Multiprocess mode enabled")
    }

    var mapWorker worker.Worker = worker.MapWorker{ TestField : "mapper"}
    var reduceWorker worker.Worker = worker.ReduceWorker{ TestField: "reducer" }

    mapWorker.Execute("fakefile")
    reduceWorker.Execute("fakefile")

    standardCoordinator := worker.NewStandardWorkerCoordinator()
    standardCoordinator.RegisterWorker(mapWorker)
    standardCoordinator.RegisterWorker(reduceWorker)
    standardCoordinator.PrintLists()

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

    partitionedFiles := freshDataPartitioner.RetrieveInputFiles()

    for _, val := range partitionedFiles {
        standardCoordinator.RegisterInputFile(val, worker.Mapper)
    }

    standardCoordinator.MapReduce(worker.NewMapReduceInput("fakePath", "fakeInput", "fakeOutput"))

}
