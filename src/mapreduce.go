package main

import (
	"fmt"
	"mapreduce/cli"
	"mapreduce/log"
	mapper "mapreduce/map"
	"mapreduce/partition"
	"mapreduce/reduce"
	"mapreduce/worker"
	"os"
)

func main() {
    for _, val := range os.Args {
        fmt.Println(val)
    }
    var options cli.CLIOptions = cli.ParseCommandLineInput()
    if (options.MultiProcess) {
        fmt.Println("Multiprocess mode enabled")
    }

    var mapWorker worker.Worker = worker.NewMapWorker("mapper-1", "intermediate-output")
    var reduceWorker worker.Worker = worker.NewReduceWorker("reducer-1", "reduce-output")

    standardCoordinator := worker.NewStandardWorkerCoordinator()
    standardCoordinator.RegisterWorker(mapWorker)
    standardCoordinator.RegisterWorker(reduceWorker)
    standardCoordinator.RegisterWorker(worker.NewMapWorker("mapper-2", "intermediate-output"))
    standardCoordinator.RegisterWorker(worker.NewReduceWorker("reducer-2", "reduce-output"))
    standardCoordinator.RegisterWorker(worker.NewMapWorker("mapper-3", "intermediate-output"))
    standardCoordinator.RegisterWorker(worker.NewReduceWorker("reducer-3", "reduce-output"))
    standardCoordinator.RegisterWorker(worker.NewMapWorker("mapper-4", "intermediate-output"))
    standardCoordinator.RegisterWorker(worker.NewReduceWorker("reducer-4", "reduce-output"))
    standardCoordinator.PrintLists()

    os.Setenv("MAPREDUCE_LOG_DEBUG_ENABLED", "enabled")
    logger := log.InitializeLog(log.LogOptions{DebugEnabled: options.LogDebugEnabled})
    logger.Debug("This should only print if debug is enabled!")
    logger.Info("This should always print!")
    logger.Debug("testing out the new sequential byte array")

    freshDataPartitioner := partition.NewSequentialDataPartitioner(20, "\n")
    freshDataPartitioner.PartitionInput(options.InputFileName)

    partitionedFiles := freshDataPartitioner.RetrieveInputFiles()

    for _, val := range partitionedFiles {
        standardCoordinator.RegisterInputFile(val, worker.Mapper)
    }

    standardCoordinator.RegisterMapper(&mapper.CountMapper{})
    standardCoordinator.RegisterReducer(&reduce.CountReducer{})

    standardCoordinator.MapReduce()

}
