package main
import (
"fmt"
    "os"
)
func main() {
    fmt.Println("hello world")
    for _, val := range os.Args {
        fmt.Println(val)
    }
}
