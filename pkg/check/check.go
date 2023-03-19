package check

import (
	"strings"

	"diplom/pkg/data"
	"diplom/pkg/storage"
)

func CheckAndFix(content []string) *storage.SMS {
	// Формирование мапу с кодами Alpha-2 по стандарту ISO 3166-1
	sourceCodes := data.Codes{}
	countries := data.CreateMap(sourceCodes)
	// Формируем мапу с именами провайдеров
	sourceProviders := data.Providers{}
	providers := data.CreateMap(sourceProviders)

	s := storage.New()

	for _, line := range content {
		field := strings.Split(line, ";")

		if len(field) != 4 {
			continue
		}

		country := field[0]
		bandwidth := field[1]
		responseTime := field[2]
		provider := field[3]

		if _, ok := countries[country]; !ok {
			continue
		}
		if _, ok := providers[provider]; !ok {
			continue
		}

		newMes := storage.NewSMS(country, bandwidth, responseTime, provider)
		s.Put(&newMes)
	}

	return &s

}
