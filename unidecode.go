package unidecode

import (
	"strings"
	"unicode"

	"github.com/mozillazg/go-unidecode/table"
)

// Version return version
func Version() string {
	return "0.1.0"
}

// Unidecode
func Unidecode(s string) string {
	return unidecode(s)
}

func unidecode(s string) string {
	runes := []rune(s)
	ret := []string{}
	for _, r := range runes {
		if r < unicode.MaxASCII {
			v := string(r)
			ret = append(ret, v)
			continue
		}
		if r > 0xeffff {
			continue
		}

		section := r >> 8   // Chop off the last two hex digits
		position := r % 256 // Last two hex digits

		if table, ok := table.Tables[section]; ok {
			if len(table) > int(position) {
				ret = append(ret, table[position])
			}
		}
	}
	return strings.Join(ret, "")
}
