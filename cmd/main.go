package main

import (
	"fmt"

	"diplom/pkg/simulator"
	"diplom/pkg/storage"
)

func main() {

	// *** 1 ***
	// Получим данные из файлов в симуляторе

	simulator := simulator.New()
	contentSms := simulator.ReadFile("sms.data")

	s := storage.New()
	contentNew := s.Get()

	// sms := check.SMS{}
	s.CheckAndFix(contentSms)

	fmt.Println("Исходный файд")
	fmt.Println()
	fmt.Println(string(contentSms))
	fmt.Println()
	fmt.Println("Обработанный")
	fmt.Println()
	for _, v := range *contentNew {
		fmt.Println(v)
	}
}
