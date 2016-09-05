package amm

import (
	"os"
	"testing"
)

var (
	testPath = "/tmp/aaa_test"
)

func TestPWD(t *testing.T) {
	cwd := pwd()
	if len(cwd) < 1 {
		t.Fatal("Current directory path empty")
	}
}

func TestExist(t *testing.T) {
	f, err := os.Create(testPath)
	if err != nil {
		t.Fatal(err)
	}

	f.Close()

	r := exist(testPath)
	if !r {
		t.Fatal("Cannot find exist file")
	}

	err = os.Remove(testPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	err := create(testPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(testPath)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(testPath)
	if err != nil {
		t.Fatal(err)
	}
}
