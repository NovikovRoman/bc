package bc

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

const (
	CryptoType PaymentSystemType = iota
	EmoneyType
	ExchangeBalanceType
	InternetBankingType
	TransferType
	CashType
)

type PaymentSystemType int

type PaymentSystem struct {
	ID         int
	Code       string
	CurrencyID int
	Type       PaymentSystemType
	Name       string
	NameAlt    string
	// [paymentSystemID]success
	Transfers map[int]bool
}

func (p PaymentSystem) Transfer(paymentSystemID int) bool {
	return p.Transfers[paymentSystemID]
}

type PaymentSystems []PaymentSystem

func NewPaymentSystems(dstDir string) (paymentSystem PaymentSystems, err error) {
	var b []byte
	if b, err = getBytes(dstDir, PaymentSystemsFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("\n"))
	paymentSystem = make(PaymentSystems, len(rows))
	// взаимосвязи переводов
	related := make([][]byte, len(rows))
	for i, row := range rows {
		cols := bytes.Split(row, []byte(";"))
		if len(cols) != 7 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		var e error
		p := PaymentSystem{
			Transfers: map[int]bool{},
		}
		if p.ID, e = strconv.Atoi(string(cols[0])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d id %v", i, e))
			continue
		}

		var idx int
		if idx, e = strconv.Atoi(string(cols[1])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d idx %v", i, e))
			continue
		}
		idx--

		p.Name = string(cols[2])
		p.NameAlt = string(cols[3])

		if p.CurrencyID, e = strconv.Atoi(string(cols[4])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d currencyID %v", i, e))
			continue
		}

		var t int
		if t, e = strconv.Atoi(string(cols[5])); e != nil {
			err = errors.Join(err, fmt.Errorf("row %d exchangeType %v", i, e))
			continue
		}
		p.Type = PaymentSystemType(t)
		if p.Type > CashType || p.Type < CryptoType {
			err = errors.Join(err, fmt.Errorf("row %d unknown exchangeType %d", i, t))
			continue
		}

		related[idx] = cols[6]
		paymentSystem[idx] = p
	}

	for psIdx, row := range related {
		for idx, b := range row {
			paymentSystem[psIdx].Transfers[idx] = b == '1'
		}
	}
	return
}
