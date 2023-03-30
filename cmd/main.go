package main

import (
	"instalmentPayments/pkg/service"
	"instalmentPayments/pkg/types"
)

func main() {
	s := &service.Service{}

	period := types.TwelveMonth
	category := types.Smartphone
	amount := 500.0
	phone := "+992935598877"

	sum, splitInMonth, err := s.GetSumOfInstalment(category, amount, phone, period)
	if err != nil {
		return
	}

	message, err := s.CreateMessageTextForInstalment(sum, splitInMonth, period)
	if err != nil {
		return
	}

	err = s.SendMessage(phone, message)
	if err != nil {
		return
	}
}
