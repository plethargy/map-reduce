package cli

import (
    "flag"
)

type CLIOptions struct {
    MultiProcess bool
    LogDebugEnabled bool
    InputFileName string
}

func ParseCommandLineInput() CLIOptions {
    multiProcessPtr := flag.Bool("multiprocess", false, "Switch between multiprocess and single-process multithread operation mode")
    logDebug := flag.Bool("debug", false, "Enables debug logging")
    fileName := flag.String("file", "stressTest.txt", "File to process")
    flag.Parse()   
    return CLIOptions{MultiProcess: *multiProcessPtr, LogDebugEnabled: *logDebug, InputFileName: *fileName}
}


