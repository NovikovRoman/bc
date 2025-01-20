package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type TopPaymentSystem struct {
	FromPaymentSystemID int
	ToPaymentSystemID   int
	Percent             float64
}

// from id exchange, to id exchange, percent
func NewTopPaymentSystems(dstDir string) (t []TopPaymentSystem, err error) {
	var b []byte
	if b, err = getBytes(dstDir, TopPaymentSystemsFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("\n"))
	t = make([]TopPaymentSystem, len(rows))
	for i, row := range rows {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 3 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		var e error
		if t[i].FromPaymentSystemID, e = strconv.Atoi(string(cols[0])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d from exchange id %v", i, e))
		}

		if t[i].ToPaymentSystemID, e = strconv.Atoi(string(cols[1])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d to exchange id %v", i, e))
		}

		if t[i].Percent, e = strconv.ParseFloat(string(cols[2]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d Reserve %v", i, e))
		}
	}
	return
}
