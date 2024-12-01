package lib

import (
	"path/filepath"
	"runtime"
)

// GetInputContent returns the path to a file next to the calling file.
func GetInputContent() (string, error) {
	// Hack to build a relative path based on the stack frame
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	dayDir := filepath.Base(dir)
	filePath := filepath.Join(dir, "..", "..", "private", "inputs", "2024", dayDir+".txt")

	return LoadFileContent(filePath)
}
