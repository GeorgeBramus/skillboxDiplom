package main

import (
	"fmt"

	"diplom/pkg/check"
	"diplom/pkg/simulator"
)

func main() {

	simulator := simulator.New()
	contentSms := simulator.ReadFile("sms.data")

	newContentSms := check.CheckAndFix(contentSms)

	for _, sms := range contentSms {
		fmt.Println(sms)
	}

	fmt.Println()

	for _, sms := range *newContentSms {
		fmt.Println(*sms)
	}

}
