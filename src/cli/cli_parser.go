package cli

import (
    "flag"
)

type CLIOptions struct {
    MultiProcess bool
}

func ParseCommandLineInput() CLIOptions {
    multiProcessPtr := flag.Bool("multiprocess", false, "Switch between multiprocess and single-process multithread operation mode")
    flag.Parse()   
    return CLIOptions{*multiProcessPtr}
}


