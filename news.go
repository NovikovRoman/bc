package bc

import (
	"bytes"
	"errors"
	"fmt"
)

type News struct {
	Title   string
	Date    string
	Content string
}

func NewNews(dstDir string) (n []News, err error) {
	var b []byte
	if b, err = getBytes(dstDir, NewsFile); err != nil {
		return
	}

	rows := bytes.Split(b, []byte("[entry_del]"))
	n = make([]News, len(rows))
	for i, row := range rows {
		cols := bytes.Split(row, []byte("[value_del]"))
		if len(cols) != 3 {
			err = errors.Join(err, fmt.Errorf("row %d has %d columns %v", i, len(cols), cols))
			continue
		}

		n[i].Title = string(cols[0])
		n[i].Content = string(cols[1])
		n[i].Date = string(cols[2])
	}
	return
}
