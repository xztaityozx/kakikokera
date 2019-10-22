package conv

import "fmt"
import "strings"

const (
	Kaki   = "柿"
	Kokera = "杮"
)

// Encode は文字列を柿と杮にエンコードする
func Encode(str string) string {
	rt := ""
	for _, v := range []rune(str) {
		rt += strings.Replace(strings.Replace(fmt.Sprintf("%064b", v), "0", Kaki, -1), "1", Kokera, -1)
	}
	return rt
}
