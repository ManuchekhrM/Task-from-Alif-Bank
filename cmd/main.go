package main

import (
	"fmt"
	"instalmentPayments/pkg/service"
	"instalmentPayments/pkg/types"
	"os"
)

func main() {
	s := &service.Service{}

	for {
		var period string
		var category string
		var amount float64
		var phone string
		fmt.Print("Добро пожаловать в калькулятор подсчета рассрочки!\n")

		fmt.Printf("\nВведите категорию из списка: %s, %s и %s:\n", types.Smartphone, types.Laptop, types.TV)
		fmt.Fscan(os.Stdin, &category)

		fmt.Printf("Введитe периюд рассрочки из списка: %s, %s, %s, %s, %s и %s месяца:\n",
			types.ThreeMonth, types.SixMonth, types.NineMonth, types.TwelveMonth, types.EighteenMonth, types.TwentyFourMonth)
		fmt.Fscan(os.Stdin, &period)

		fmt.Print("Введите сумму выплаты. К приверу 100.0:\n")
		fmt.Fscan(os.Stdin, &amount)

		fmt.Print("Введите номер телефона для отправки информации о рассрочке:\n")
		fmt.Fscan(os.Stdin, &phone)

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
}
