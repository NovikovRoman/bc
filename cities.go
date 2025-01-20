package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type City struct {
	ID   int
	Name string
}

func NewCities(dstDir string) (c []City, err error) {
	var b []byte
	if b, err = getBytes(dstDir, CitiesFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("\n"))
	c = make([]City, len(rows))
	for i, row := range rows {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 2 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		var e error
		if c[i].ID, e = strconv.Atoi(string(cols[0])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, e))
		}

		c[i].Name = string(cols[1])
	}
	return
}
