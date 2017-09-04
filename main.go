package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/phayes/checkstyle"
)

var (
	cstyle = checkstyle.New()
)

var (
	lintMinConfidence = flag.Float64("lint_min_confidence", 0.8, "lint: minimum confidence of a problem to print it")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgolint [flags] # runs on package in current directory\n")
	fmt.Fprintf(os.Stderr, "\tgolint [flags] [packages]\n")
	fmt.Fprintf(os.Stderr, "\tgolint [flags] [directories] # where a '/...' suffix includes all sub-directories\n")
	fmt.Fprintf(os.Stderr, "\tgolint [flags] [files] # all must belong to a single package\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Do linting
	lintDir(".")

	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Println(cstyle)
}
