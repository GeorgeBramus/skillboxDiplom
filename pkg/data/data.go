package data

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

type Alpha2 struct {
	Code    string `csv:"code"`
	Country string `csv:"country"`
}

func GetAlphaCodes() map[string]string {
	iso, err := os.OpenFile("../data/standarts/ISO_3166-1.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer iso.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	codeAlpha := []*Alpha2{}

	if err := gocsv.UnmarshalFile(iso, &codeAlpha); err != nil {
		panic(err)
	}

	codMap := make(map[string]string)
	for _, code := range codeAlpha {
		codMap[code.Code] = code.Country
	}

	return codMap
}
