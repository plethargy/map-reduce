package main
import (
    "fmt"
    "os"
    "mapreduce/cli"
    "mapreduce/fileHandler"
    "mapreduce/worker"
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

    success := fileHandler.FileBasedOutputStream{}.OutputData("Hello world to file! From interface implementation", "newFile.txt")
    if success {
        fmt.Println("Successfully wrote to file")
    } else {
        fmt.Println("Failed to write to file")
    }
    var mapWorker worker.Worker = worker.MapWorker{ TestField : "mapper"}
    var reduceWorker worker.Worker = worker.ReduceWorker{ TestField: "reducer" }

    mapWorker.Execute()
    reduceWorker.Execute()

    standardCoordinator := worker.StandardWorkerCoordinator{}
    standardCoordinator.RegisterWorker(mapWorker)
    standardCoordinator.RegisterWorker(reduceWorker)
    standardCoordinator.PrintLists()
}
