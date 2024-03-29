package log

import (
    "fmt"
    "os"
    "strings"
)

type Logger interface {
    Debug(output string)
    Info(output string)
    Error(output string)
}

func ListEnvironmentVariables() {
    fmt.Println("Listing environment variables")
    for _, env := range os.Environ() {
        pair := strings.SplitN(env, "=", 2)
        key := pair[0]
        val := pair[1]
        fmt.Printf("%s=%s\n", key, val)
        
    }
}
type StandardLogger struct {
}

type LogOptions struct {
    DebugEnabled bool
}

func (sl StandardLogger) Debug(output string) {
    if os.Getenv("MAPREDUCE_LOG_DEBUG_ENABLED") == "enabled" {
        fmt.Println(output)
    }
}

func (sl StandardLogger) Info(output string) {
    fmt.Println(output)
}

func (sl StandardLogger) Error(output string) {
    fmt.Println(output) //TODO: make this actually warn on output
}

func InitializeLog(logOptions LogOptions) Logger {
    os.Unsetenv("MAPREDUCE_LOG_DEBUG_ENABLED")
    if logOptions.DebugEnabled {
        os.Setenv("MAPREDUCE_LOG_DEBUG_ENABLED", "enabled")
    }
    return StandardLogger{}
}

type LogHandler struct {
    logger Logger
}

func (lh LogHandler) Debug(output string) {
    lh.logger.Debug(output)
}
func (lh LogHandler) Info(output string) {
    lh.logger.Info(output)
}
func (lh LogHandler) Error(output string) {
    lh.logger.Error(output)
}
func NewLogHandler(logger Logger) LogHandler {
    return LogHandler{logger: logger}
}


