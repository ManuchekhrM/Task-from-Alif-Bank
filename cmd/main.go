package main

import (
	"fmt"
	"instalmentPayments/pkg/service"
	"instalmentPayments/pkg/types"
	"strings"
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

	sumForMsg := fmt.Sprintf("%.2f", sum)
	splitInMon := fmt.Sprintf("%.2f", splitInMonth)

	msg := "Вы совершили покупку с рассрочкой на {period} месяцев на {sum} сомони. Ваша ежемесячная оплата составляет {per month} сомон."

	message := strings.ReplaceAll(msg, "{period}", period)
	message = strings.ReplaceAll(message, "{sum}", sumForMsg)
	message = strings.ReplaceAll(message, "{per month}", splitInMon)

	err = s.SendMessage(phone, message)
	if err != nil {
		return
	}
}
