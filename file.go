package amm

import (
	"encoding/json"
	"errors"
)

var (
	// ErrExist ...
	ErrExist = errors.New("metadata.json exist")
)

// Init ...
func Init() error {
	mpath := mpath()
	r := exist(mpath)
	if !r {
		m := &Metadata{}
		data, err := json.MarshalIndent(m, "", "\t")
		if err != nil {
			return err
		}
		err = create(data, mpath)
		if err != nil {
			return err
		}
	}
	if r {
		return ErrExist
	}

	return nil
}
