package service

import (
	"errors"
	"fmt"
	"instalmentPayments/pkg/types"
	"log"
	"math"
	"strconv"
)

var (
	ErrPhoneRegistered      = errors.New("phone already registered")
	ErrAmountMustBePositive = errors.New("amount must be greater")
	ErrInvalidPeriod       = errors.New("ErrInvalidPeriod")
)

type Service struct {
}

// Gives us percent for installments for additional months
func Percents(category string, period string) float64 {
	switch category {
	case types.Smartphone:
		if period == types.TwelveMonth {
			periodPercent := 3.0
			return periodPercent
		}
		if period == types.EighteenMonth {
			periodPercent := 6.0
			return periodPercent
		}
		if period == types.TwentyFourMonth {
			periodPercent := 9.0
			return periodPercent
		}
	case types.Laptop:
		if period == types.EighteenMonth {
			periodPercent := 4.0
			return periodPercent
		}
		if period == types.TwentyFourMonth {
			periodPercent := 8.0
			return periodPercent
		}
	case types.TV:
		if period == types.TwentyFourMonth {
			periodPercent := 5.0
			return periodPercent
		}
	default:
		log.Panicln(ErrInvalidPeriod)
	}
	return 0
}

// We calculate the instalment amount with percent
func (s *Service) GetSumOfInstalment(category string, amount float64, phone string, period string) (sum float64, splitInMonth float64, err error) {
	if amount <= 0 {
		return 0, 0, ErrAmountMustBePositive
	}

	percentOfCategory := Percents(category, period)
	sum = ((100.0 + percentOfCategory) * amount) / 100.0

	per, _ := strconv.ParseFloat(period, 32)
	splitInMonth = sum / per

	// Rounded to 2 digits
	splitInMonth = math.Ceil((splitInMonth)*100) / 100

	return sum, splitInMonth, nil
}

// Send info message about instalment
func (s *Service) SendMessage(phone string, message string) error {
	// Тут должно было быть подключено СМС центр для отправки сообщения
	fmt.Println(message)
	return nil
}
