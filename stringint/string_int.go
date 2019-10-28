package stringint

import (
	"errors"
)

// StringToInt ...
func StringToInt(s string) (int, error) {
	if len(s) == 0 || len(s) > 10 {
		return 0, errors.New("param error")
	}

	p := s[0]
	if p == '-' {
		s = s[1:]
	}

	bs := []byte(s)
	var n int
	for _, v := range bs {
		ch := v - '0'
		if ch < 0 || ch > 9 {
			return 0, errors.New("param error invail " + string(v))
		}
		n = n*10 + int(ch)
	}
	if p == '-' {
		n = n * -1
	}
	return n, nil
}

// IntToString ...
func IntToString(i int) string {
	px := ""
	if i < 0 {
		px = "-"
		i = i * -1
	}
	table := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	s := ""
	for {
		if i < 10 {
			s = string(table[i]) + s
			return px + s
		}

		p := i % 10
		s = string(table[p]) + s
		i = i / 10
	}
}
