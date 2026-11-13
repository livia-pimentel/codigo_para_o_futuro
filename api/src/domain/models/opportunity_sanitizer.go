package models

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Remove espaços excedentes
func CleanString(s string) string {
	// Remove espaços excedentes do meio e das pontas
	cleaned := strings.Join(strings.Fields(s), " ")
	if cleaned == "" {
		return ""
	}

	// Transforma letras iniciais em maiúscula e demais em minúsculas
	c := cases.Title(language.Portuguese)
	return c.String(cleaned)
}