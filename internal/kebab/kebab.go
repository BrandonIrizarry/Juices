package kebab

import "strings"

func KebabCase(name string) string {
	subwords := strings.Fields(name)

	for i := range subwords {
		subwords[i] = strings.ToLower(subwords[i])
	}

	return strings.Join(subwords, "-")
}

func UndoKebabCase(itemName string) string {
	parts := strings.Split(itemName, "-")

	for i := range parts {
		subword := parts[i]
		parts[i] = strings.ToUpper(string(subword[0])) + string(subword[1:])
	}

	return strings.Join(parts, " ")
}
