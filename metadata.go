package amm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Metadata ...
type Metadata struct {
	Metadata Anime `json:"metadata"`
	path     string
}

// Exist check metadata.json exist
func (m *Metadata) Exist() error {
	cwd := CWD()
	m.path = filepath.Join(cwd, "metadata.json")

	src, err := os.Stat(m.path)
	if err != nil {
		return err
	}
	if src.IsDir() {
		return errors.New("metadata.json is directory")
	}

	return nil
}

// Create create new metadata.json
func (m *Metadata) Create() error {
	data, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(m.path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
