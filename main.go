package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"slices"
)

const (
	Reset = "\033[0m"
	Blue  = "\033[34m"
)

func getCurrentWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error occured while getting cwd: %v", err)
	}

	return dir
}

func openCurrentWorkingDir(currentWorkingDir string) []fs.DirEntry {
	dir, err := os.ReadDir(currentWorkingDir)
	if err != nil {
		log.Fatalf("Error opening current working directory: %v", err)
	}

	return dir
}

func main() {
	var err error
	dir := getCurrentWorkingDir()

	longListFlag := flag.Bool("l", false, "use a long listing format")
	allListFlag := flag.Bool("a", false, "do not ignore entries starting with .")
	humanReadableFlag := flag.Bool("h", false, "with -l and -s, print sizes like 1K 234M 2G etc.")
	listOneFilePerLine := flag.Bool(
		"1",
		false,
		"list one file per line.  Avoid '\\n' with -q or -b",
	)

	// Preprocess flags to support combined flags like -la
	args := preprocessArgs(os.Args[1:])

	flag.CommandLine.Parse(args)

	if flag.NFlag() == 0 ||
		(containsAny(args, []string{"-a", "-h", "-1"}) && !slices.Contains(args, "-l")) {
		ListFiles(dir, *allListFlag, *listOneFilePerLine)
	}

	switch *longListFlag {
	case true:
		err = LongListFiles(dir, *allListFlag, *humanReadableFlag)
	}

	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(0)
	}
}
