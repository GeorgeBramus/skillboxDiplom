package simulator

import (
	"io"
	"io/ioutil"
	"os"

	"diplom/pkg/logs"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logs.InitialSet("main")
}

type Simulator struct {
	Directory string
}

// New - создание Симулятора. Установим директорию по умолчанию.
func New() Simulator {
	viper := viper.New()
	viper.SetConfigFile("../data/config/config.yml")
	viper.ReadInConfig()

	directory := viper.GetString("simulator.directory")

	s := Simulator{
		Directory: directory,
	}
	return s
}

// ReadFile - чтение файла в директории симулятора
// и возврат содержимого указанного файла.
func (s *Simulator) ReadFile(name string) []byte {
	f := s.Directory + name
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadFile(f)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	return content
}
