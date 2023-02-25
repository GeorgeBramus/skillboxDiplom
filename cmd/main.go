package main

import (
	"fmt"

	"diplom/pkg/data"
	"diplom/pkg/simulator"
)

func main() {
	// Данные из симулятора для sms-сервиса
	simulator := simulator.New()
	contentSms := simulator.ReadFile("sms.data")

	// Формирование мапы с кодами Alpha-2 по стандарту ISO 3166-1
	// [?] Двубуквенные коды для обозначения страны
	codes := data.GetAlphaCodes()

	fmt.Println(codes["RU"])
	fmt.Println(string(contentSms))
}
