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
		err = m.Write()
		if err != nil {
			l.Println(err)
			return
		}

		l.Println("metadata.json created")
		return
	}

	l.Println(errExist)
}

// Add ...
func Add(args []string) {
	m := &Metadata{}

	err := m.Exist()
	if err != nil {
		l.Fatal(err)
	}

	err = m.Load()
	if err != nil {
		l.Fatal(err)
	}

	switch args[0] {
	case "title":
		t := NewTitle(args[1:])
		m.Metadata.Title = *t

	case "atitle":
		t := NewTitle(args[1:])
		m.Metadata.AlternativeTitles = append(m.Metadata.AlternativeTitles, *t)

	default:
		return
	}

	err = m.Write()
	if err != nil {
		l.Fatal(err)
	}
}
