package utils

import (
	"errors"
	"strings"
	"unicode"
)

func ValidateNickname(nickname string) error {
	if len(nickname) == 0 {
		return errors.New("nickname cannot be empty")
	}

	if len(nickname) > 20 {
		return errors.New("nickname cannot have more than 20 characters")
	}

	for _, r := range nickname {
		if unicode.IsSpace(r) {
			return errors.New("nickname cannot contain spaces")
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return errors.New("the nickname can only contain letters, numbers and underscores")
		}
	}

	return nil
}

// Verifica se uma mensagem é válida (não vazia e dentro de limites de tamanho)
func ValidateMessage(message string) error {
	if len(strings.TrimSpace(message)) == 0 {
		return errors.New("message cannot be empty")
	}

	if len(message) > 500 {
		return errors.New("the message cannot have more than 500 characters")
	}

	return nil
}
