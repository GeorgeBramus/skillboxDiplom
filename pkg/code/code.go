package code

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// Parse Парсинг файла с таблицей кодов ISO 3166-1 Alpha-2
func Parse() {
	// code := Create()

	file, err := os.Open("../data/raw/ISO_3166-1.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file2, err2 := os.OpenFile("../data/ISO_3166-1.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err2 != nil {
		panic(err2)
	}
	defer file2.Close()

	fileScanner := bufio.NewScanner(file)
	reName := regexp.MustCompile(`>[А-Яа-я\s]+<`)
	reAlpha := regexp.MustCompile(`>[A-Z]{2}<`)
	var textName, textAlpha string

	for fileScanner.Scan() {
		textName = clear(reName.FindString(fileScanner.Text()))
		textAlpha = clear(reAlpha.FindString(fileScanner.Text()))

		if textName != "" {
			t := textAlpha + ";" + textName + "\n"
			_, err := file2.WriteString(t)
			if err != nil {
				panic(err)
			}
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

// ***
// Функции в помощь

func clear(s string) string {
	str := strings.Replace(s, "<", "", 1)
	str = strings.Replace(str, ">", "", 1)
	return str
}
