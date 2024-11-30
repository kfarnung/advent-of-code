package lib

import (
	"path/filepath"
	"runtime"
)

// GetTestFilePath returns the path to a file next to the calling file.
func GetTestFilePath(name string) string {
	// Hack to build a relative path based on the stack frame
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	dayDir := filepath.Base(dir)
	return filepath.Join(dir, "..", "..", "private", "inputs", "2020", dayDir+".txt")
}
