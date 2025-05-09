package validationtype

import (
	"strings"
	"unicode"
)

type Password struct {
	Value               string
	message             string
	hasMinLen           bool
	hasUpper            bool
	hasLower            bool
	hasNumber           bool
	hasSpecialCharacter bool
}

func NewPassword(value string) *Password {
	return &Password{Value: value}
}

func (p *Password) Message() string {
	return p.message
}

func (p *Password) Validate() bool {
	const minLen = 8
	if len(p.Value) >= minLen {
		p.hasMinLen = true
	}

	for _, char := range p.Value {
		switch {
		case unicode.IsUpper(char):
			p.hasUpper = true
		case unicode.IsLower(char):
			p.hasLower = true
		case unicode.IsNumber(char):
			p.hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			p.hasSpecialCharacter = true
		}
	}

	if !p.hasMinLen || !p.hasUpper || !p.hasLower || !p.hasNumber || !p.hasSpecialCharacter {
		p.setMessage()
		return false
	}

	return true
}

func (p *Password) setMessage() {
	var s strings.Builder

	if !p.hasMinLen {
		s.WriteString("min 8 characters, ")
	}
	if !p.hasUpper {
		s.WriteString("uppercase character required, ")
	}
	if !p.hasLower {
		s.WriteString("lowercase character required, ")
	}
	if !p.hasNumber {
		s.WriteString("numeric character required, ")
	}
	if !p.hasSpecialCharacter {
		s.WriteString("special character required, ")
	}

	p.message = strings.Trim(s.String(), ", ")
}
