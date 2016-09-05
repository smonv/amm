package amm

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var (
	l = log.New(os.Stdout, "", 0)

	errExist = errors.New("metadata.json exist")
)

// Init ...
func Init() {
	mpath := mpath()
	r := exist(mpath)
	if !r {
		m := &Metadata{}
		data, err := json.MarshalIndent(m, "", "    ")
		if err != nil {
			l.Fatal(err)
		}
		err = create(data, mpath)
		if err != nil {
			l.Fatal(err)
		}
	}
	if r {
		l.Fatal(errExist)
	}

	l.Println("metadata.json created")
}
