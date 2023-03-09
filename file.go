package go_utils

import (
	"io/fs"
	"log"
	"os"
	"path"
)

func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	} else {
		return !os.IsNotExist(err)
	}
}

func TryToRead(files []string) ([]byte, error) {
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err == nil {
			return content, err
		} else {
			if !os.IsNotExist(err) {
				log.Printf("read %s fail: %s", file, err)
			}
		}
	}
	return nil, fs.ErrNotExist
}

func LocateFile(name string) ([]byte, error) {
	targets := make([]string, 0)

	wd, err := os.Getwd()
	if err == nil && wd != "" {
		targets = append(targets, path.Join(wd, name))
	}

	homeDir, err := os.UserHomeDir()
	if err == nil && homeDir != "" {
		targets = append(targets, path.Join(homeDir, name))
	}

	executable, err := os.Executable()
	if err == nil {
		binDir := path.Dir(executable)
		if binDir != "" {
			targets = append(targets, path.Join(binDir, name))
		}
	}

	return TryToRead(targets)
}
