package service

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

func ValidateDist(text string) (float64, error) {
	text = strings.TrimSpace(text)
	s := strings.ReplaceAll(text, ",", ".")

	dist, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1.0, errors.New("Не верный ввод, введи число в километрах\nТоесть без букв, только цифры")
	}

	if dist < 0 {
		return -1, errors.New("Братан, как ты мог пробежать отрицательное число километров?")
	}

	if dist == 0 {
		return -1, errors.New("Я не буду это считать")
	}

	if dist > 200 {
		return -1, errors.New("Ага, а до Сатурна ты не добежал случайно?\nЛибо напоминаю, что ввод в километрах, а не метрах")
	}

	return dist, nil
}

func ValidateName(text string) (string, error) {
	name := strings.TrimSpace(text)

	if name == "" {
		return "", errors.New("Имя не должно быть пустым, введи ещё раз")
	}

	if utf8.RuneCountInString(name) < 3 {
		return "", errors.New("Имя не может быть короче 3-х символов")
	}

	return name, nil
}
