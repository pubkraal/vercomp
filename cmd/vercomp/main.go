package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pubkraal/vercomp/internal/version"
)

func flagUsage() {
	fmt.Printf("Usage: %s version version [version...]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = flagUsage
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// TODO move code below this to a package
	v1 := flag.Arg(0)
	v2 := flag.Arg(1)

	ver1, err := version.New(v1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid version input: %v", v1)
		os.Exit(127)
	}
	ver2, err := version.New(v2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid version input: %v", v2)
		os.Exit(127)
	}

	if v1 == v2 {
		os.Exit(0)
	}

	if ver1.Less(ver2) {
		os.Exit(-1)
	} else {
		os.Exit(1)
	}
}
