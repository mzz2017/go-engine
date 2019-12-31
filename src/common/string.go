package common

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

const gcharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func IntArrayToString(a []int, delim string) string {
	var ret string
	for _, s := range a {
		ret += strconv.Itoa(s) + delim
	}
	return ret
}

func Int32ArrayToString(a []int32, delim string) string {
	var ret string
	for _, s := range a {
		ret += strconv.Itoa((int)(s)) + delim
	}
	return ret
}

func Int64ArrayToString(a []int64, delim string) string {
	var ret string
	for _, s := range a {
		ret += strconv.Itoa((int)(s)) + delim
	}
	return ret
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func RandStr(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = gcharset[RandInt31n(len(gcharset))]
	}
	return string(b)
}

type StrTableLine struct {
	cols []string
}

type StrTable struct {
	header []string
	lines  []StrTableLine
}

func (s *StrTable) AddHeader(h string) {
	s.header = append(s.header, h)
}

func (s *StrTable) AddLine(l StrTableLine) {
	s.lines = append(s.lines, l)
}

func (s *StrTableLine) AddData(d string) {
	s.cols = append(s.cols, d)
}

func (s *StrTable) String() string {

	if len(s.header) <= 0 {
		return ""
	}

	colmax := make([]int, 0)
	for _, s := range s.header {
		colmax = append(colmax, len(s))
	}

	totalcol := 0
	for i := 0; i < len(colmax); i++ {
		max := colmax[i]
		for _, sl := range s.lines {
			if i < len(sl.cols) {
				max = MaxOfInt(max, len(sl.cols[i]))
			}
		}
		colmax[i] = max
		totalcol += max
	}
	totalcol += len(colmax) + 1

	/*
		-----------
		| a  | b  |
		-----------
		| 1  | 2  |
		-----------
	*/

	ret := ""
	ret += strings.Repeat("-", totalcol) + "\n"
	for i, h := range s.header {
		ret += "|" + WrapString(h, colmax[i])
	}
	ret += "|" + "\n"

	for _, l := range s.lines {
		ret += strings.Repeat("-", totalcol) + "\n"
		for i, d := range l.cols {
			ret += "|" + WrapString(d, colmax[i])
		}
		for i := len(l.cols); i < len(colmax); i++ {
			ret += "|" + WrapString("", colmax[i])
		}
		ret += "|" + "\n"
	}

	ret += strings.Repeat("-", totalcol) + "\n"

	return ret
}

func WrapString(s string, n int) string {
	if n <= len(s) {
		return s
	}
	l := (n - len(s)) / 2
	r := (n - len(s)) - l
	return strings.Repeat(" ", l) + s + strings.Repeat(" ", r)
}
