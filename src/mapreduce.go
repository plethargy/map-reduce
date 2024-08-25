package main

import (
	"fmt"
	"mapreduce/cli"
	"mapreduce/log"
	"mapreduce/map"
	"mapreduce/partition"
	"mapreduce/reduce"
	"mapreduce/worker"
	"encoding/json"
	"os"
)

const configFileName string = "config.json"

type Configuration struct {
    MapWorkerCount int
    ReduceWorkerCount int
    MapWorkerOutputPath string
    ReduceWorkerOutputPath string
}

func main() {
    for _, val := range os.Args {
        fmt.Println(val)
    }
    file, _ := os.Open(configFileName)
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := Configuration{}
    err := decoder.Decode(&config)
    if err != nil {
	println("Unable to read %s, using default config instead", configFileName)
	config = NewDefaultConfiguration()
    }

    println("Using %d map workers and %d reduce workers", config.MapWorkerCount, config.ReduceWorkerCount)
    
    var options cli.CLIOptions = cli.ParseCommandLineInput()
    if (options.MultiProcess) {
        fmt.Println("Multiprocess mode enabled")
    }

    standardCoordinator := worker.NewStandardWorkerCoordinator()

    for i := 0; i < config.MapWorkerCount; i++ {
	mapWorker := worker.NewMapWorker(fmt.Sprintf("mapper-%d", i + 1), config.MapWorkerOutputPath)
	standardCoordinator.RegisterWorker(mapWorker)
    }

    for i := 0; i < config.ReduceWorkerCount; i++ {
	reduceWorker := worker.NewReduceWorker(fmt.Sprintf("reducer-%d", i + 1), config.ReduceWorkerOutputPath)
	standardCoordinator.RegisterWorker(reduceWorker)
    }
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

func NewDefaultConfiguration() Configuration {
    return Configuration{MapWorkerCount: 1, ReduceWorkerCount: 1, MapWorkerOutputPath: "default-intermediate-output", ReduceWorkerOutputPath: "default-reduce-output"}
}
