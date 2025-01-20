package bc

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	LastUpdate        time.Time
	CurrentVersion    string
	CompatibleVersion string
}

func NewInfo(dstDir string) (i Info, err error) {
	var b []byte
	if b, err = getBytes(dstDir, InfoFile); err != nil {
		return
	}

	for _, row := range bytes.Split(b, []byte("\n")) {
		cols := bytes.Split(row, []byte("="))
		if len(cols) != 2 {
			continue
		}

		switch string(cols[0]) {
		case "last_update":
			i.LastUpdate = parseLastUpdate(cols[1])

		case "current_version":
			i.CurrentVersion = string(cols[1])

		case "compatible_version":
			i.CompatibleVersion = string(cols[1])
		}
	}
	return
}

var months = map[string]time.Month{
	"января":   1,
	"февраля":  2,
	"марта":    3,
	"апреля":   4,
	"мая":      5,
	"июня":     6,
	"июля":     7,
	"августа":  8,
	"сентября": 9,
	"октября":  10,
	"ноября":   11,
	"декабря":  12,
}

func parseLastUpdate(b []byte) (lu time.Time) {
	now := time.Now()
	loc, _ := time.LoadLocation("Europe/Moscow")
	// "17:14:58, 4 июля"
	m := regexp.MustCompile(`(\d{2}:\d{2}:\d{2}),\s*(\d{1,2})\s*([^\s]+)`).FindSubmatch(b)
	if len(m) == 0 {
		return
	}

	month, ok := months[strings.ToLower(string(m[3]))]
	if !ok {
		return
	}

	t, _ := time.Parse("15:04:05", string(m[1]))
	day, _ := strconv.Atoi(string(m[2]))
	return time.Date(now.Year(), month, day, t.Hour(), t.Minute(), t.Second(), 0, loc)
}
