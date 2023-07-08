package address

import "strings"

func contains(slice []string, word string) bool {
	for _, v := range slice {
		if v == word {
			return true
		}
	}
	return false
}

func AddressType(address string) string {
	addresses := []string{"rua", "avenida", "estrada", "rodovia"}

	firstWord := strings.ToLower(address[:strings.Index(address, " ")])

	if contains(addresses, firstWord) {
		return strings.Title(firstWord)
	}

	return "TIPO INVÁLIDO"
}
