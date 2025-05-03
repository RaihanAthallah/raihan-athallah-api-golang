package helper

import (
	"strings"
	"unicode"
)

func GenerateSlug(title string) string {
	// Ubah ke huruf kecil
	slug := strings.ToLower(title)
	// Ganti spasi dengan -
	slug = strings.ReplaceAll(slug, " ", "-")
	// Hapus karakter non-alphanumeric
	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			return r
		}
		return -1
	}, slug)
	return slug
}
