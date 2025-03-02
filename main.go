package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookMarks := bookmarkMap{}
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			checkBookMarks(bookMarks)
		case 2:
			addNewBookMark(bookMarks)
		case 3:
			deleteBookMark(bookMarks)
		case 4:
			break Menu
		}
	}
}

// Просмотреть закладки
func checkBookMarks(bookMarks bookmarkMap) {
	if len(bookMarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookMarks {
		fmt.Println(key, ": ", value)
	}
}

// Добавить закладку
func addNewBookMark(toMap bookmarkMap) {
	key := ""
	value := ""

	fmt.Printf("Введите название закладки: ")
	fmt.Scanln(&key)
	fmt.Printf("Введите сайт для закладки: \"%s\":\n", key)
	fmt.Scanln(&value)
	toMap[key] = value
	fmt.Printf("Закладка \"%s\" добавлена!\n", key)
}

// Удалить закладку
func deleteBookMark(fromMap bookmarkMap) {
	key := ""

	fmt.Printf("Введите название закладки для удаления")
	fmt.Scanln(&key)
	delete(fromMap, key)
	fmt.Printf("Закладка \"%s\" удалена!\n", key)
}

// Вызов Меню
func getMenu() int {
	fmt.Println("***___Утилита закладки___***\n - 1. Посмотреть закладки\n - 2. Добавить закладку\n - 3. Удалить закладку \n - 4. Выход")
	var userChoise int
	fmt.Scanln(&userChoise)

	return userChoise
}
