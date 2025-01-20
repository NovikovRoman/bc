package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ExchangeRate struct {
	CurrencyIDFrom int
	CurrencyIDTo   int
	ExchangeID     int
	AmountFrom     float64
	AmountTo       float64
	Reserve        float64
	Pretensions    int
	Reviews        int
	Active         bool
	Min            float64
	Max            float64
	CityID         int
}

func NewExchangeRates(dstDir string) (r []ExchangeRate, err error) {
	var b []byte
	if b, err = getBytes(dstDir, ExchangeRatesFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("\n"))
	r = make([]ExchangeRate, len(rows))
	for i, row := range bytes.Split(b, []byte("\n")) {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 11 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		var e error
		exRate := ExchangeRate{}
		if exRate.CurrencyIDFrom, e = strconv.Atoi(string(cols[0])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id currencyFrom %v", i, e))
			continue
		}
		if exRate.CurrencyIDTo, e = strconv.Atoi(string(cols[1])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id currencyTo %v", i, e))
			continue
		}
		if exRate.ExchangeID, e = strconv.Atoi(string(cols[2])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id exchange %v", i, e))
			continue
		}
		if exRate.AmountFrom, e = strconv.ParseFloat(string(cols[3]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d amount from %v", i, e))
			continue
		}
		if exRate.AmountTo, e = strconv.ParseFloat(string(cols[4]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d amount to %v", i, e))
			continue
		}
		if exRate.Reserve, e = strconv.ParseFloat(string(cols[5]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d reserve %v", i, e))
			continue
		}

		pretensionsReviews := strings.Split(string(cols[6]), ".")
		if len(pretensionsReviews) != 2 {
			err = errors.Join(err, fmt.Errorf("row %d bad pretensionsReviews %s", i, cols[6]))
			continue
		}
		if exRate.Pretensions, e = strconv.Atoi(pretensionsReviews[0]); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d pretensions %v", i, e))
			continue
		}
		if exRate.Reviews, e = strconv.Atoi(pretensionsReviews[1]); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d reviews %v", i, e))
			continue
		}

		exRate.Active = string(cols[7]) == "1"

		if exRate.Min, e = strconv.ParseFloat(string(cols[8]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d min %v", i, e))
			continue
		}
		if exRate.Max, e = strconv.ParseFloat(string(cols[9]), 64); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d max %v", i, e))
			continue
		}
		if exRate.CityID, e = strconv.Atoi(string(cols[10])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id city %v", i, e))
			continue
		}

		r[i] = exRate
	}
	return
}
