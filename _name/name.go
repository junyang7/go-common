package _name

import (
	"strings"
	"unicode"
)

func UpperCamelCase(name string) string {
	words := splitWords(name)
	for i := 0; i < len(words); i++ {
		words[i] = capitalize(words[i])
	}
	return strings.Join(words, "")
}
func LowerCamelCase(name string) string {
	words := splitWords(name)
	if len(words) == 0 {
		return ""
	}
	for i := 1; i < len(words); i++ {
		words[i] = capitalize(words[i])
	}
	words[0] = strings.ToLower(words[0])
	return strings.Join(words, "")
}
func SnakeCase(name string) string {
	words := splitWords(name)
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "_")
}

func splitWords(name string) []string {
	var words []string
	var buf []rune
	for i, r := range name {
		if r == '_' || r == '-' || r == ' ' {
			if len(buf) > 0 {
				words = append(words, string(buf))
				buf = buf[:0]
			}
			continue
		}
		if i > 0 && unicode.IsUpper(r) && len(buf) > 0 {
			words = append(words, string(buf))
			buf = buf[:0]
		}
		buf = append(buf, unicode.ToLower(r))
	}
	if len(buf) > 0 {
		words = append(words, string(buf))
	}
	return words
}
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
