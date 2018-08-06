package main

import (
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// UserHome returns the current user home path
func UserHome() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}

// FileExists checks if the provided file exists or not
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Warnf("File does not exist: %s", filename)
		return false
	}

	log.Debugf("File exists: %s", filename)
	return true
}
