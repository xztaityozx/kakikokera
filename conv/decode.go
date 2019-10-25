package conv

import (
	"strconv"

	"golang.org/x/xerrors"
)

// Decode は文字列strから柿と杮を取り出して0と1に置き換え、8つに区切って文字列へ変換する
func Decode(str string) (string, error) {
	kaki := []rune(Kaki)[0]
	kokera := []rune(Kokera)[0]

	var box []byte
	bits := ""

	for _, v := range str {
		if v == kaki {
			bits += "0"
		} else if v == kokera {
			bits += "1"
		}

		if len(bits) == 8 {
			b, err := strconv.ParseUint(bits, 2, 8)
			if err != nil {
				return "", err
			}
			box = append(box, byte(b))
			bits = ""
		}
	}

	if len(box) == 0 {
		return "", xerrors.New("デコードできる文字列が空だが")
	}

	return string(box), nil
}
