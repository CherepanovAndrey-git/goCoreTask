package pkg

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

/*
StrConv принимает на себя срез с произвольным кол-вом значений типа interface{},
преобразует каждое значение из среза в строку, конкатенирует и возвращает объединенную строку.
*/
func StrConv(vars ...interface{}) string {
	var builder strings.Builder
	for _, v := range vars {
		builder.WriteString(fmt.Sprintf("%v", v))
	}
	return builder.String()
}

// AddSaltAndHash добавляет соль в середину среза рун от строки и возвращает SHA256 хэш
func AddSaltAndHash(s string) string {
	saltedStr := AddSaltToRunes(s)
	hash := sha256.Sum256([]byte(saltedStr))
	return fmt.Sprintf("%x", hash)
}

// AddSaltToRunes добавляет соль в середину строки аргумента
func AddSaltToRunes(s string) string {
	salt := "go-2024"
	runes := []rune(s)
	mid := len(runes) / 2
	saltedRunes := append(runes[:mid], append([]rune(salt), runes[mid:]...)...)
	return string(saltedRunes)
}
