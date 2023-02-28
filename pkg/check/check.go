package check

import (
	"bufio"
	"strings"
)

type SMSData struct {
	Сountry      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func checkSms(rawData []string, countFields int, codes, providers map[string]string) []SMSData {
	var (
		fields []string
		data   []SMSData
	)
	for _, line := range rawData {
		fields = strings.Split(line, ";")
		if len(fields) != countFields {
			continue
		}
		if _, exists := codes[fields[0]]; !exists {
			continue
		}
		if _, exists := providers[fields[3]]; !exists {
			continue
		}

		country := fields[0]

		sms := SMSData{
			Country:      country,
			Bandwidth:    fields[1],
			ResponseTime: fields[2],
			Provider:     fields[3],
		}

		data = append(data)
	}
}
