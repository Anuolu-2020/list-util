package main

import (
	"fmt"
	"strconv"
	"strings"
)

const width = 4

// For ls command
func ListFiles(currentWorkingDir string, printAll bool, listPerLine bool) {
	dir := openCurrentWorkingDir(currentWorkingDir)

	for _, dirContent := range dir {
		file, _ := dirContent.Info()

		// Skip dot files
		if !printAll && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		// indicate a directory with color
		if file.IsDir() {
			fmt.Printf(Blue+"%s"+Reset, file.Name())
			continue
		}

		if listPerLine {
			fmt.Printf("%s\n", file.Name())
		} else {
			fmt.Printf("%s\t", file.Name())
		}
	}

	fmt.Println()
}

// For ls -l command
func LongListFiles(currentWorkingDir string, printAll bool, humanReadable bool) error {
	dir := openCurrentWorkingDir(currentWorkingDir)
	totalDiskBlock := getTotalDiskBlock(dir, false)

	if humanReadable {
		size, err := strconv.ParseFloat(totalDiskBlock, 64)
		if err != nil {
			return err
		}
		totalDiskBlock = FormatSize(size)
	}

	fmt.Printf("total: %s\n", totalDiskBlock)
	for _, dirContent := range dir {
		file, _ := dirContent.Info()

		// Skip dot files
		if !printAll && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fileLinkCount := getFileLinkCount(file)
		owner, group, err := getFileOwnerAndGroup(file)
		if err != nil {
			return err
		}
		modTime := file.ModTime().Format("Jan 2 15:04")

		// indicate a directory with color
		if file.IsDir() {
			var fileSize string

			if humanReadable {
				fileSize = FormatSize(float64(file.Size()))
			} else {
				fileSize = fmt.Sprint(file.Size())
			}

			fmt.Printf(
				"%-10s %*d %*s %*s %*s %*s "+Blue+"%*s\n"+Reset,
				file.Mode().Perm().String(),
				width,
				fileLinkCount,
				width,
				owner,
				width,
				group,
				width,
				fileSize,
				width,
				modTime,
				width,
				file.Name(),
			)
			continue
		}

		var fileSize string

		if humanReadable {
			fileSize = FormatSize(float64(file.Size()))
		} else {
			fileSize = fmt.Sprint(file.Size())
		}

		fmt.Printf(
			"%-10s %*d %*s %*s %*s %*s %s\n",
			file.Mode().Perm().String(),
			width,
			fileLinkCount,
			width,
			owner,
			width,
			group,
			width,
			fileSize,
			width,
			modTime,
			file.Name(),
		)
	}

	fmt.Println()

	return nil
}
