package utils

import "strings"

// Pluralize returns the plural form of a given English noun.
func Pluralize(word string) string {
	irregularPlurals := map[string]string{
		"child":  "children",
		"goose":  "geese",
		"man":    "men",
		"woman":  "women",
		"foot":   "feet",
		"tooth":  "teeth",
		"go":     "goes",
		"person": "people",
	}

	if plural, ok := irregularPlurals[word]; ok {
		return plural
	}

	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") ||
		strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	if strings.HasSuffix(word, "y") {
		return strings.TrimSuffix(word, "y") + "ies"
	}

	return word + "s"
}
