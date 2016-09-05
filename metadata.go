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

// Write create new metadata.json
func (m *Metadata) Write() error {
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

// Load ...
func (m *Metadata) Load() error {
	data, err := ioutil.ReadFile(m.path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}
