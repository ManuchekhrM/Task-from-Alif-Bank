package main

import (
	"instalmentPayments/pkg/service"
	"log"
)

func main() {
	s := &service.Service{}

	for {
		period, category, amount, phone, err := s.AddFromConsole()
		if err != nil {
			log.Printf("ERR: %v \n\n", err)
			continue
		}

		sum, splitInMonth, err := s.GetSumOfInstalment(category, amount, phone, period)
		if err != nil {
			log.Printf("ERR: %v \n\n", err)
			continue
		}

		message, err := s.CreateMessageTextForInstalment(sum, splitInMonth, period)
		if err != nil {
			log.Printf("ERR: %v \n\n", err)
			continue
		}

		err = s.SendMessage(phone, message)
		if err != nil {
			log.Printf("ERR: %v \n\n", err)
			continue
		}
	}
}
