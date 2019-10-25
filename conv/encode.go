package conv

import (
	"fmt"
	"strings"

	"golang.org/x/xerrors"
)

const (
	Kaki   = "柿"
	Kokera = "杮"
)

// Encode は文字列を柿と杮にエンコードする
func Encode(str string) (string, error) {

	if len(str) == 0 {
		return "", xerrors.New("入力文字列が空ですが")
	}

	rt := ""
	for _, v := range []byte(str) {
		rt += strings.Replace(strings.Replace(fmt.Sprintf("%08b", v), "0", Kaki, -1), "1", Kokera, -1)
	}
	return rt, nil
}
