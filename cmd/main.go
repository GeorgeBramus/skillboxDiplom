package main

import (
	"fmt"
	"strings"

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

	p := "Gmail"
	providers := data.GetNameProviders()
	if strings.Contains(providers[p], ",") {
		services := strings.Split(providers[p], ",")
		fmt.Println("Один сервис:", services[0])
		fmt.Println("Другой сервис:", services[1])
	} else {
		fmt.Println("Этот оператор относится только к одному сервису:", providers[p])
	}

	fmt.Printf("\n\n")
	fmt.Println(codes["RU"])
	fmt.Println(string(contentSms))
}
