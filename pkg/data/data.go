package data

import (
	"bufio"
	"os"
	"strings"

	"diplom/pkg/logs"

	log "github.com/sirupsen/logrus"
)

type Codes struct{}
type Providers struct{}

type SourceData interface {
	FileName() string
}

func init() {
	logs.InitialSet("main")
}

// CreateMap - Возвращает мапу с сырыми данными
func CreateMap(s SourceData) map[string]string {
	file, err := os.Open(s.FileName())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	codMap := make(map[string]string)

	var lines []string
	for fileScanner.Scan() {
		lines = strings.Split(fileScanner.Text(), ";")
		codMap[lines[0]] = lines[1]
	}

	return codMap
}

// Codes.FileName - Возвращает имя файла, в котором сырые данные
func (c Codes) FileName() string {
	return "../data/standarts/ISO_3166-1.csv"
}

// Providers.FileName - Возвращает имя файла, в котором сырые данные
func (p Providers) FileName() string {
	return "../data/standarts/providers.csv"
}
