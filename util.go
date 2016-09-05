package amm

import "os"

// CWD return current working directory
func CWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return cwd
}
