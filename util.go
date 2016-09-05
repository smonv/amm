package amm

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return wd
}

func mpath() string {
	cwd := pwd()
	if len(cwd) < 1 {
		return ""
	}

	mp := filepath.Join(cwd, "metadata.json")
	return mp
}

func exist(path string) bool {
	src, err := os.Stat(path)
	if err != nil {
		return false
	}

	if src.IsDir() {
		return false
	}

	return true
}

func create(data []byte, path string) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
