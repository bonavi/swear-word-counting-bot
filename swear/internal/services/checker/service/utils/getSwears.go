package utils

import "strings"

// Функция для проверки наличия ругательства в сообщении
func GetSwears(message string, swearsMap map[string]struct{}) []string {

	var swears []string

	// Разбиваем сообщение на слова
	words := strings.Fields(message)

	// Проходимся по каждому слову
	for _, word := range words {

		// Приводим каждое слово к нижнему регистру
		word = strings.ToLower(word)

		// Если слово найдено в словаре
		if _, found := swearsMap[word]; found {
			swears = append(swears, word)
		}
	}

	return swears
}
