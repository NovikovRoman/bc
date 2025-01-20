package bc

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"
)

const (
	CurrenciesFile         = "bm_bcodes.dat"
	CurrencyRatesFile      = "bm_brates.dat"
	CitiesFile             = "bm_cities.dat"
	PaymentSystemsFile     = "bm_cy.dat"
	PaymentSystemCodesFile = "bm_cycodes.dat"
	ExchagesFile           = "bm_exch.dat"
	InfoFile               = "bm_info.dat"
	NewsFile               = "bm_news.dat"
	ExchangeRatesFile      = "bm_rates.dat"
	TopPaymentSystemsFile  = "bm_top.dat"
)

func getBytes(dstDir, name string) (b []byte, err error) {
	var f *os.File
	if f, err = os.Open(filepath.Join(dstDir, name)); err != nil {
		return
	}
	defer f.Close()

	return io.ReadAll(charmap.Windows1251.NewDecoder().Reader(f))
}
