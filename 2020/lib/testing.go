package lib

import (
	"path/filepath"
	"runtime"
)

// GetInputFilePath returns the path to the input file.
func GetInputFilePath() string {
	// Hack to build a relative path based on the stack frame
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	dayDir := filepath.Base(dir)
	return filepath.Join(dir, "..", "..", "private", "inputs", "2020", dayDir+".txt")
}
