package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type Currency struct {
	ID   int
	Code string
	Name string
	Bank string
}

func NewCurrencies(dstDir string) (c []Currency, err error) {
	var b []byte
	if b, err = getBytes(dstDir, CurrenciesFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("\n"))
	c = make([]Currency, len(rows))
	for i, row := range rows {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 4 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		var e error
		if c[i].ID, e = strconv.Atoi(string(cols[0])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, e))
		}

		c[i].Code = string(cols[1])
		c[i].Name = string(cols[2])
		c[i].Bank = string(cols[3])
	}
	return
}
