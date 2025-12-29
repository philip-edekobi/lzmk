package utils

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

const MaxSlugLen = 72

// Slugify converts a title into a filesystem- and URL-safe slug.
// Behavior is deterministic and MUST NOT change once shipped.
func Slugify(title string) (string, error) {
	if title == "" {
		return "", fmt.Errorf("Empty title detected")
	}

	// Normalize to NFKD (decompose accents)
	normed := norm.NFKD.String(title)

	var b strings.Builder
	b.Grow(len(normed))

	dash := false

	for _, r := range normed {
		// Strip diacritics
		if unicode.Is(unicode.Mn, r) {
			continue
		}

		r = unicode.ToLower(r)

		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			dash = false

		case r >= '0' && r <= '9':
			b.WriteRune(r)
			dash = false

		default:
			if !dash && b.Len() > 0 {
				b.WriteByte('-')
				dash = true
			}
		}

		if b.Len() >= MaxSlugLen {
			break
		}
	}

	slug := strings.Trim(b.String(), "-")

	if slug == "" {
		return "", fmt.Errorf("Problem with generating slug")
	}

	return slug, nil
}
