package cli

import (
    "flag"
)

type CLIOptions struct {
    MultiProcess bool
    LogDebugEnabled bool
}

func ParseCommandLineInput() CLIOptions {
    multiProcessPtr := flag.Bool("multiprocess", false, "Switch between multiprocess and single-process multithread operation mode")
    logDebug := flag.Bool("debug", false, "Enables debug logging")
    flag.Parse()   
    return CLIOptions{MultiProcess: *multiProcessPtr, LogDebugEnabled: *logDebug}
}


