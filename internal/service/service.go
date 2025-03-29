package service

import (
	"errors"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

type MorseService interface {
	AutoConvert(input string) (string, error) // Переименовали для ясности
	Encode(text string) (string, error)       // Явное кодирование
	Decode(morse string) (string, error)      // Явное декодирование
}

type morseServiceImpl struct {
	converter *morse.Converter
}

func NewMorseService() MorseService {
	return &morseServiceImpl{
		converter: &morse.DefaultConverter,
	}
}

// DetectAndConvert автоматически определяет тип и конвертирует
func (s *morseServiceImpl) AutoConvert(input string) (string, error) {
	if input == "" {
		return "", ErrEmptyInput
	}

	if isTextInput(input) {
		return s.Encode(input) // Текст → Morse
	}
	return s.Decode(input) // Morse → Текст
}

// Encode явно кодирует текст в Morse
func (s *morseServiceImpl) Encode(text string) (string, error) {
	if text == "" {
		return "", ErrEmptyInput
	}
	return s.converter.ToMorse(text), nil
}

// Decode явно декодирует Morse в текст
func (s *morseServiceImpl) Decode(morse string) (string, error) {
	if morse == "" {
		return "", ErrEmptyInput
	}
	return s.converter.ToText(morse), nil
}

// isTextInput определяет, является ли ввод текстом (true) или Morse-кодом (false)
func isTextInput(input string) bool {
	hasLettersOrNumbers := false
	hasNonMorseChars := false

	for _, r := range input {
		switch {
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			hasLettersOrNumbers = true
		case r != '.' && r != '-' && r != ' ' && !unicode.IsSpace(r):
			hasNonMorseChars = true
		}
	}

	return hasLettersOrNumbers || hasNonMorseChars
}

// ErrEmptyInput to show error
var ErrEmptyInput = errors.New("empty input")
