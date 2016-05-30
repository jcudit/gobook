package main

import (
	"path/filepath"
	"strings"
)

func CommonPrefix(paths []string) string {
	numPaths := len(paths)
	for i := 0; i < minLength(paths); i++ {
		for j := 1; j < numPaths; j++ {
			if paths[j][i] != paths[0][i] {
				if i == 0 {
					return ""
				} else {
					return paths[0][:i]
				}
			}
		}
	}
	return paths[0]
}

func minLength(paths []string) int {
	var shortest int
	for i, path := range paths {
		if i == 0 || len(path) < shortest {
			shortest = len(path)
		}
	}
	return shortest
}

func CommonPathPrefix(paths []string) string {
	if len(paths) < 2 {
		return ""
	}
	components := make([][]string, len(paths))

	// Convert each path into a slice of path components
	for i, path := range paths {
		components[i] = strings.SplitAfter(path, string(filepath.Separator))
		if len(components[i]) == 1 {
			return ""
		}
	}
	// Find index where components stop matching
	endIndex := commonUntil(components)

	// Return the joined common prefix path
	return filepath.Join(components[0][:endIndex]...)
}

func commonUntil(components [][]string) int {
	result := 0
	// Compare first row to subsequent rows
Outer:
	for col := range components[0] { // 0 .. numCol
		for row := 1; row < len(components); row++ { // 1 .. numRow
			if len(components[row]) >= col && components[0][col] == components[row][col] {
				continue
			} else {
				// Return column where rows differ
				break Outer
			}
		}
		result++
	}
	return result
}
