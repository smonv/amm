package amm

import (
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
	m := &Metadata{}

	err := m.Exist()
	if err != nil {
		err = m.Create()
		if err != nil {
			l.Println(err)
			return
		}

		l.Println("metadata.json created")
		return
	}

	l.Println(errExist)
}
