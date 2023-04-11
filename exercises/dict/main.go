package main

import "fmt"

func main() {
	// створення словника англійських слів з їх українським перекладом
	dict := make(map[string]string)
	dict["apple"] = "яблуко"
	dict["banana"] = "банан"
	dict["cherry"] = "вишня"
	dict["grape"] = "виноград"
	dict["lemon"] = "лимон"

	// запит перекладу певного слова
	word := "apple"
	translation, ok := dict[word]
	if ok {
		fmt.Printf("%s перекладається як %s\n", word, translation)
	} else {
		fmt.Printf("Переклад слова %s не знайдено\n", word)
	}
}
