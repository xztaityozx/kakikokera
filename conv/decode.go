package conv

import (
	"strconv"
	"strings"
)

// Decode は文字列strから柿と杮を取り出して0と1に置き換え、8つに区切って文字列へ変換する
func Decode(str string) string {
	kaki := []rune(Kaki)[0]
	kokera := []rune(Kokera)[0]

	target := ""

	for _, v := range str {
		if v == kaki || v == kokera {
			target += string(v)
		}
	}
	bits := strings.Replace(strings.Replace(target, Kaki, "0", -1), Kokera, "1", -1)
	var bin string
	var box []rune
	for _, v := range bits {
		bin += string(v)
		if len(bin) == 64 {
			res, _ := strconv.ParseInt(bin, 2, 64)
			box = append(box, rune(res))
			bin = ""
		}
	}
	return string(box)
}
