package golang_string_title

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func ToTitle(str string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(str); i++ {
		if unicode.IsLetter(rune(str[i])) {
			if i == 0 {
				buffer.WriteRune(unicode.ToTitle(rune(str[i])))
			} else if i > 0 && unicode.IsSpace(rune(str[i-1])) {
				buffer.WriteRune(unicode.ToTitle(rune(str[i])))
			}else {
				buffer.WriteRune(rune(str[i]))
			}
		} else {
			buffer.WriteRune(rune(str[i]))
		}
	}
	return buffer.String()
}

func both(s string) {
	fmt.Printf("[%s] -> [%v] [%v] [%v]\n", s, strings.ToTitle(s), strings.Title(s), ToTitle(s))
}

func main() {
	fmt.Println("[Original] -> [strings.ToTitle] [strings.Title] [zstrings.ToTitle]")
	both("s")
	both("s h")
	both("s h m")
	both("stulle horst")
	both("s h2\tm\nf")
	// Those three panic in zstrings.ToTitle()
	both("s  h")
	both("")
	both("  s  h  ")
	both("uncle jim's supermarket")
}
