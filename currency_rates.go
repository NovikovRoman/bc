package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type CurrencyRates map[int]map[int]float64

func NewCurrencyRates(dstDir string) (r CurrencyRates, err error) {
	var b []byte
	if b, err = getBytes(dstDir, CurrencyRatesFile); err != nil {
		return
	}

	r = CurrencyRates{}

	for i, row := range bytes.Split(b, []byte("\n")) {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 3 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		id1, e := strconv.Atoi(string(cols[0]))
		if e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, e))
			continue
		}

		var id2 int
		if id2, e = strconv.Atoi(string(cols[1])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, e))
			continue
		}

		if _, ok := r[id1]; !ok {
			r[id1] = map[int]float64{}
		}

		if r[id1][id2], e = strconv.ParseFloat(string(cols[2]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d rate %v", i, e))
			continue
		}

		if _, ok := r[id2]; !ok {
			r[id2] = map[int]float64{}
		}
		if _, ok := r[id2][id1]; !ok {
			r[id2][id1] = 1 / r[id1][id2]
		}
	}
	return
}
