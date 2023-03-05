package check

import (
	"bufio"
	"diplom/pkg/data"
	"diplom/pkg/storage"
	"strings"
)

func (s *storage.SMS) CheckAndFix(content []byte) {
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio.ScanLines)

	// Формирование мапы с кодами Alpha-2 по стандарту ISO 3166-1
	sourceCodes := data.Codes{}
	countries := data.CreateMap(sourceCodes)
	// Формируем мапу с именами провайдеров
	sourceProviders := data.Providers{}
	providers := data.CreateMap(sourceProviders)

	var line []string

	for scanner.Scan() {
		line = strings.Split(scanner.Text(), ";")
		country := line[0]
		bandwidth := line[1]
		responseTime := line[2]
		provider := line[3]

		if _, ok := countries[country]; !ok {
			continue
		}
		if _, ok := providers[provider]; !ok {
			continue
		}

		newMes := storage.NewSMS(country, bandwidth, responseTime, provider)
		s := storage.New()
		s.Put(&newMes)
	}
}
