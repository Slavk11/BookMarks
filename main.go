package main

import "fmt"

func main() {
	bookMarks := map[string]string{}

	for {
		fmt.Println("***___Утилита закладки___***\n - 1. Посмотреть закладки\n - 2. Добавить закладку\n - 3. Удалить закладку \n - 4. Выход")
		var userChoise int
		fmt.Scanln(&userChoise)

		switch userChoise {
		case 1:
			checkBookMarks(bookMarks)
		case 2:
			key := ""
			value := ""

			fmt.Printf("Введите название закладки: ")
			fmt.Scanln(&key)
			fmt.Printf("Введите сайт для закладки: \"%s\":\n", key)
			fmt.Scanln(&value)
			addNewBookMark(key, value, bookMarks)
			fmt.Printf("Закладка \"%s\" добавлена!\n", key)
		case 3:
			key := ""
			fmt.Printf("Введите название закладки для удаления")
			deleteBookMark(key, bookMarks)
			fmt.Printf("Закладка \"%s\" удалена!\n", key)

		case 4:
			return
		}
	}
}

// Просмотреть закладки
func checkBookMarks(bookMarks map[string]string) {
	fmt.Println(bookMarks)
}

// Добавить закладку
func addNewBookMark(key string, value string, toMap map[string]string) map[string]string {
	toMap[key] = value
	return toMap
}

// Удалить закладку
func deleteBookMark(key string, fromMap map[string]string) map[string]string {
	fmt.Scanln(&key)
	delete(fromMap, key)
	return fromMap
}
