package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <entryPoints.json>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	//todo process inputfile

	done := make(chan bool)
	//todo long task
	<-done
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
