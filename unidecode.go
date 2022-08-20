package unidecode

import (
	"strings"
	"unicode"

	"github.com/mozillazg/go-unidecode/table"
)

// Version return version
func Version() string {
	return "0.2.0"
}

// Unidecode implements transliterate Unicode text into plain 7-bit ASCII.
// e.g. Unidecode("kožušček") => "kozuscek"
func Unidecode(s string) string {
	return unidecode(s)
}

func unidecode(s string) string {
	var ret strings.Builder
	for _, r := range s {
		if r < unicode.MaxASCII {
			ret.WriteRune(r)
			continue
		}
		if r > 0xeffff {
			continue
		}

		section := r >> 8   // Chop off the last two hex digits
		position := r % 256 // Last two hex digits
		if tb, ok := table.Tables[section]; ok {
			if len(tb) > int(position) {
				ret.WriteString(tb[position])
			}
		}
	}
	return ret.String()
}
