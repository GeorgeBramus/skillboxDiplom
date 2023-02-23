package main

import (
	"fmt"

	"diplom/pkg/simulator"
)

func main() {
	simulator := simulator.New()
	contentSms := simulator.ReadFile("sms.data")

	fmt.Printf("%s\n", contentSms)
}
