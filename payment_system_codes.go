package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type PaymentSystemCodes map[int]string

func NewPaymentSystemCodes(dstDir string) (codes PaymentSystemCodes, err error) {
	var b []byte
	if b, err = getBytes(dstDir, PaymentSystemCodesFile); err != nil {
		return
	}

	codes = PaymentSystemCodes{}
	rows := bytes.Split(b, []byte("\n"))
	for i, row := range rows {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 2 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		id, e := strconv.Atoi(string(cols[0]))
		if e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, err))
			continue
		}

		codes[id] = string(cols[1])
	}
	return
}
