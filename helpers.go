package main

import (
	"fmt"
	"io/fs"
	"math"
	"os/user"
	"slices"
	"strings"
	"syscall"
)

func containsAny(slice []string, items []string) bool {
	for _, item := range items {
		if slices.Contains(slice, item) {
			return true
		}
	}
	return false
}

// preprocessArgs splits combined flags like -la into separate flags (-l -a)
func preprocessArgs(args []string) []string {
	var processedArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 2 {
			// Break combined flags (e.g., -la) into separate ones (-l -a)
			for _, char := range arg[1:] {
				processedArgs = append(processedArgs, fmt.Sprintf("-%c", char))
			}
		} else {
			processedArgs = append(processedArgs, arg)
		}
	}
	return processedArgs
}

func getFileStat(file fs.FileInfo) *syscall.Stat_t {
	if sys := file.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			return stat
		}
	}

	return nil
}

func getFileOwnerAndGroup(file fs.FileInfo) (string, string, error) {
	uid := uint32(0)
	gid := uint32(0)

	stat := getFileStat(file)
	uid = stat.Uid
	gid = stat.Gid

	owner, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		return "", "", err
	}

	group, err := user.LookupGroupId(fmt.Sprint(gid))
	if err != nil {
		return "", "", err
	}

	return owner.Username, group.Name, nil
}

func getFileLinkCount(file fs.FileInfo) uint64 {
	nlink := uint64(0)
	stat := getFileStat(file)
	nlink = uint64(stat.Nlink)

	return nlink
}

func getTotalDiskBlock(dir []fs.DirEntry, countHiddenFiles bool) string {
	totalDiskBlock := uint64(0)
	for _, dirContent := range dir {
		file, _ := dirContent.Info()

		if strings.HasPrefix(file.Name(), ".") && !countHiddenFiles {
			continue
		}

		if sys := file.Sys(); sys != nil {
			if stat, ok := sys.(*syscall.Stat_t); ok {
				totalDiskBlock += uint64(stat.Blocks)
			}
		}
	}

	return fmt.Sprint(totalDiskBlock)
}

func FormatSize(size float64) string {
	var dividedSize float64
	if size < 1024 {
		return fmt.Sprintf("%.0fB", size)
	}

	dividedSize = size / 1024

	if dividedSize < 1024 {
		if dividedSize == math.Trunc(dividedSize) {
			return fmt.Sprintf("%.0fK", dividedSize)
		} else {
			return fmt.Sprintf("%.1fK", dividedSize)
		}
	}

	dividedSize /= 1024

	if dividedSize < 1024 {
		if dividedSize == math.Trunc(dividedSize) {
			return fmt.Sprintf("%.0fM", dividedSize)
		} else {
			return fmt.Sprintf("%.1fM", dividedSize)
		}
	}

	dividedSize /= 1024

	if dividedSize < 1024 {
		if dividedSize == math.Trunc(dividedSize) {
			return fmt.Sprintf("%.0fG", dividedSize)
		} else {
			return fmt.Sprintf("%.1fG", dividedSize)
		}
	}

	dividedSize /= 1024

	if dividedSize < 1024 {
		if dividedSize == math.Trunc(dividedSize) {
			return fmt.Sprintf("%.0fT", dividedSize)
		} else {
			return fmt.Sprintf("%.1fT", dividedSize)
		}
	}

	return ""
}
