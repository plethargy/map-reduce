package main
import (
    "fmt"
    "os"
    "mapreduce/cli"
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
}
