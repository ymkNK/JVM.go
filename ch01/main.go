package main

import "fmt"

func main() {
	cmd :=parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1 by ymk 2020")
	}else if cmd.helpFlag{
		printUsage()
	}else{
		startJVM(cmd);
	}
}

