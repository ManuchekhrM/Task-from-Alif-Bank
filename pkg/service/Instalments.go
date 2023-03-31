package service

import (
	"errors"
	"fmt"
	"instalmentPayments/pkg/types"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	ErrAmountMustBePositive = errors.New("Amount must be greater")
	ErrInvalidPeriod        = errors.New("ErrInvalidPeriod")
)

type Service struct {
}

func (s *Service) AddFromConsole() (period string, category string, amount float64, phone string, err error) {
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

	return
}

// Gives us percent for installments for additional months
func (s *Service) Percents(category string, period string) (float64, error) {
	switch category {
	case types.Smartphone:
		switch period {
		case types.TwelveMonth:
			return 3.0, nil
		case types.EighteenMonth:
			return 6.0, nil
		case types.TwentyFourMonth:
			return 9.0, nil
		default:
			return 0, nil
		}
	case types.Laptop:
		switch period {
		case types.EighteenMonth:
			return 4.0, nil
		case types.TwentyFourMonth:
			return 8.0, nil
		default:
			return 0, nil
		}
	case types.TV:
		switch period {
		case types.TwentyFourMonth:
			return 5.0, nil
		default:
			return 0, nil
		}
	default:
		return 0, ErrInvalidPeriod
	}
}

// We calculate the instalment amount with percent
func (s *Service) GetSumOfInstalment(category string, amount float64, phone string, period string) (sum float64, splitInMonth float64, err error) {
	if amount <= 0 {
		return 0, 0, ErrAmountMustBePositive
	}

	percentOfCategory, _ := s.Percents(category, period)

	sum = ((100.0 + percentOfCategory) * amount) / 100.0

	per, _ := strconv.ParseFloat(period, 32)
	splitInMonth = sum / per

	// Rounded to 2 digits
	splitInMonth = math.Ceil((splitInMonth)*100) / 100

	return sum, splitInMonth, nil
}

func (s *Service) CreateMessageTextForInstalment(sum float64, splitInMonth float64, period string) (string, error) {
	sumForMsg := fmt.Sprintf("%.2f", sum)
	splitInMon := fmt.Sprintf("%.2f", splitInMonth)

	msg := "Вы совершили покупку с рассрочкой на {period} месяцев на {sum} сомони. Ваша ежемесячная оплата составляет {per month} сомон.\n \n"

	message := strings.ReplaceAll(msg, "{period}", period)
	message = strings.ReplaceAll(message, "{sum}", sumForMsg)
	message = strings.ReplaceAll(message, "{per month}", splitInMon)

	return message, nil
}

// Send info message about instalment
func (s *Service) SendMessage(phone string, message string) error {
	// Тут должно было быть подключено СМС центр для отправки сообщения
	fmt.Println(message)
	return nil
}
