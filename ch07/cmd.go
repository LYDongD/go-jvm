package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	verboseClassFlag bool
	verboseInstFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

//parse to cmd from os arg
func parseCmd() *Cmd {
	cmd := &Cmd{}

	//if parse error, printUsage
	flag.Usage = printUsage

	//define flag and store value to cmd struct
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose outpu")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose outpu")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose outpu")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	//parse common-line flag from os.Arg[1]
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...] \n", os.Args[0])
}
