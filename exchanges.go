package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type Exchange struct {
	ID      int
	Name    string
	WMBL    int // WebMoney Business Level
	Reserve int
	// ACTS int // Advanced Cash Transaction Score
	// PMTS int // Perfect Money Trust Score
}

// id, name, ?, WMBL, Reserve
func NewExchanges(dstDir string) (ex []Exchange, err error) {
	var b []byte
	if b, err = getBytes(dstDir, ExchagesFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("\n"))
	ex = make([]Exchange, len(rows))
	for i, row := range rows {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 5 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		var e error
		if ex[i].ID, e = strconv.Atoi(string(cols[0])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, e))
		}

		ex[i].Name = string(cols[1])

		if ex[i].WMBL, e = strconv.Atoi(string(cols[3])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d WMBL %v", i, e))
		}

		if ex[i].Reserve, e = strconv.Atoi(string(cols[4])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d Reserve %v", i, e))
		}
	}
	return
}
