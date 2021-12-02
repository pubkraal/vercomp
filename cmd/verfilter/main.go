package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pubkraal/vercomp/pkg/version"
)

func flagUsage() {
	fmt.Printf("Usage: %s [OPTIONS] version version [version...]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	// FIXME add a parameter to set minimum and to print a sorted list of higher versions back
	minver := flag.String("minimum", "", "define a minimum to print versions sorted from positional arguments")
	printsep := flag.String("separator", "\n", "the separator between printed versions")
	flag.Usage = flagUsage
	flag.Parse()

	slice, err := version.NewSlice(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not parse all versions: %v", err)
		os.Exit(1)
	}

	var mv *version.Version
	if *minver != "" {
		mv, err = version.New(*minver)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse minimum version: %v", err)
			os.Exit(1)
		}
	}

	slice.Sort()

	for idx, v := range slice {
		if mv != nil && (v.Less(mv) || v.Original == mv.Original) {
			continue
		}
		fmt.Print(v.Repr())
		if idx != len(slice)-1 {
			fmt.Print(*printsep)
		}
	}

	fmt.Print("\n")
}
