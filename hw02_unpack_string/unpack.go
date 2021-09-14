package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder // https://otus.ru/nest/post/717/
	var prevSimbol string       // предидущий символ

	for _, symbol := range str {
		if unicode.IsDigit(symbol) { // нашли цифру
			if len(prevSimbol) == 0 { // первый символ цифра или несколько цифр подряд
				return "", ErrInvalidString
			}

			number, err := strconv.Atoi(string(symbol)) // привели rune к int (Repeat(s string, count int))
			// логируем в случае ошибки, но по идее ее тут быть не может
			if err != nil {
				fmt.Printf("Error strconv.Atoi %s", err)
			}
			// дописали в строку символ и повторили его number раз (0 удалит 1 символ)
			builder.WriteString(strings.Repeat(prevSimbol, number))
			prevSimbol = "" // "обнулили" предидущий чтоб "поймать" две цифры подряд
			continue
		}

		builder.WriteString(prevSimbol) // если не цифра просто пишем один раз
		prevSimbol = string(symbol)
	}

	builder.WriteString(prevSimbol) // записали последний символ, если он не цифра

	return builder.String(), nil
}
