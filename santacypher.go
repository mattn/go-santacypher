package santacypher

import (
	"strings"
)

var toHo = map[rune]string{
	'A': "HoHO",
	'B': "HOHoHoHo",
	'C': "HOHoHOHo",
	'D': "HOHoHo",
	'E': "Ho",
	'F': "HoHoHOHo",
	'G': "HOHOHo",
	'H': "HoHoHoHo",
	'I': "HoHo",
	'J': "HoHOHOHO",
	'K': "HOHoHO",
	'L': "HoHOHoHo",
	'M': "HOHO",
	'N': "HOHo",
	'O': "HOHOHO",
	'P': "HoHOHOHo",
	'Q': "HOHOHoHO",
	'R': "HoHOHo",
	'S': "HoHoHo",
	'T': "HO",
	'U': "HoHoHO",
	'V': "HoHoHoHO",
	'W': "HoHOHO",
	'X': "HOHoHoHO",
	'Y': "HOHoHOHO",
	'Z': "HOHOHoHo",
}

var fromHo = map[string]rune{
	"HoHO":     'A',
	"HOHoHoHo": 'B',
	"HOHoHOHo": 'C',
	"HOHoHo":   'D',
	"Ho":       'E',
	"HoHoHOHo": 'F',
	"HOHOHo":   'G',
	"HoHoHoHo": 'H',
	"HoHo":     'I',
	"HoHOHOHO": 'J',
	"HOHoHO":   'K',
	"HoHOHoHo": 'L',
	"HOHO":     'M',
	"HOHo":     'N',
	"HOHOHO":   'O',
	"HoHOHOHo": 'P',
	"HOHOHoHO": 'Q',
	"HoHOHo":   'R',
	"HoHoHo":   'S',
	"HO":       'T',
	"HoHoHO":   'U',
	"HoHoHoHO": 'V',
	"HoHOHO":   'W',
	"HOHoHoHO": 'X',
	"HOHoHOHO": 'Y',
	"HOHOHoHo": 'Z',
}

const breakChar = "🎅"

// encodes to ho ho ho
func ToHoHoHo(s string) string {
	s = strings.ToUpper(s)
	output := ""
	for _, r := range s {
		if s, ok := toHo[r]; ok {
			output += s + breakChar
		} else {
			output += string(r)
		}
	}
	return output
}

// decodes from hohoho
func FromHoHoHo(s string) string {
	output := ""
	for _, ss := range strings.Split(s, breakChar) {
		for len(ss) > 0 {
			if r, ok := fromHo[ss]; ok {
				output += string(r)
				ss = ""
			} else {
				rs := []rune(ss)
				output += string(rs[0])
				ss = string(rs[1:])
			}
		}
	}
	return output
}
