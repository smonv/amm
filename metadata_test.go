package amm

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	mt = &Metadata{}
)

func TestMetadataFileExist(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tp := filepath.Join(cwd, "metadata.json")

	err = ioutil.WriteFile(tp, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = mt.Exist()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(tp)
	if err != nil {
		t.Fatal(tp)
	}
}

func TestMetadataFileCreate(t *testing.T) {
	err := mt.Exist()
	if err == nil {
		t.Fatal("metadata.json exist")
	}

	err = mt.Create()
	if err != nil {
		t.Fatal(err)
	}

	err = mt.Exist()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(mt.path)
	if err != nil {
		t.Fatal(err)
	}
}
