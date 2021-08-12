package main

import (
	"os"
	"path/filepath"
	"runtime"
)

// UserHomeDir ...
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

// NormalizePath ...
func NormalizePath(path string) string {
	paths := filepath.SplitList(os.ExpandEnv(path))
	for i, i2 := range paths {
		if i2 == "~" {
			paths[i] = UserHomeDir()
		}
	}
	return filepath.Join(paths...)
}
