
import (
	"bytes"
	"strings"
)

//Crowdbotics

func ToTitle(str string) string {
	arrWords := strings.Split(str, " ")
	newString := func(str string, r rune, i int) string {
		return str[:i] + string(r) + str[i + 1:]
	}
	var buffer bytes.Buffer
	for i := 0; i < len(arrWords); i++ {
		if (arrWords[i][0] >= 'a' && arrWords[i][0] <= 'z') {
			buffer.WriteString(newString(arrWords[i], rune(arrWords[i][0]) - 32, 0))
			buffer.WriteString(" ")
		}
		else {
			buffer.WriteString(arrWords[i])
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
