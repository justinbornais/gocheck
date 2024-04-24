package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Get a string representation of a file.
// If no file type is available, returns "Unknown".
func GetFileType(path string) string {
	value := fileTypes[GetExtension(path)]
	if value == "" {
		return "Unknown"
	}
	return value
}

// Get the file extension in lowercase with the period.
//
// For example: The file "some/directory/sample.txt" will return ".txt"
func GetExtension(path string) string {
	return strings.ToLower(filepath.Ext(path))
}

// Get the file extension in lowercase without the period.
//
// For example: The file "some/directory/sample.txt" will return "txt"
func GetExtensionNoPeriod(path string) string {
	return strings.ToLower(filepath.Ext(path)[1:])
}

// Check if the path is a directory.
func IsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	return info.IsDir(), err
}

// Check if the path is a file.
func IsFile(path string) (bool, error) {
	info, err := os.Stat(path)
	return info.Mode().IsRegular(), err
}

// Get the file mode.
func GetFileMode(path string) (fs.FileMode, error) {
	info, err := os.Stat(path)
	return info.Mode(), err
}

// Get the file mode as a string.
func GetFileModeString(path string) (string, error) {
	info, err := os.Stat(path)
	return info.Mode().String(), err
}

// Get the file type.
func GetFileBitType(path string) (fs.FileMode, error) {
	info, err := os.Stat(path)
	return info.Mode().Type(), err
}

// Get the file type as a string.
func GetFileBitTypeString(path string) (string, error) {
	info, err := os.Stat(path)
	return info.Mode().Type().String(), err
}

// Get the file permissions.
func GetFilePermission(path string) (fs.FileMode, error) {
	info, err := os.Stat(path)
	return info.Mode().Perm(), err
}

// Get the file permissions as a string.
func GetFilePermissionString(path string) (string, error) {
	info, err := os.Stat(path)
	return info.Mode().Perm().String(), err
}

// Get the file size.
func GetFileSize(path string) (int, error) {
	info, err := os.Stat(path)
	return int(info.Size()), err
}

// Get the file size as an int64.
func GetFileSizeInt64(path string) (int64, error) {
	info, err := os.Stat(path)
	return info.Size(), err
}
