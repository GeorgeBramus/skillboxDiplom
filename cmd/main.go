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
	sourceCodes := data.Codes{}
	codes := data.CreateMap(sourceCodes)
	// Формируем мапу с именами провайдеров
	sourceProviders := data.Providers{}
	providers := data.CreateMap(sourceProviders)

	fmt.Println(providers["Topolo"])
	fmt.Println(codes["KZ"])

	fmt.Println()
	fmt.Println(string(contentSms))
}
