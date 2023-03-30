package service

import (
	"instalmentPayments/pkg/types"
	"reflect"
	"testing"
)

func TestGetSumOfInstalment_inTwelveMonth_forSmartphone(t *testing.T) {
	var service Service

	period := types.TwelveMonth
	category := types.Smartphone
	phone := "+992935598877"

	sum, splitInMonth, err := service.GetSumOfInstalment(category, 1000.0, phone, period)
	if err != nil {
		t.Error(err)
	}

	expectedSum := 1030.0

	if !reflect.DeepEqual(expectedSum, sum) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSum, sum)
	}

	expectedSplitInMonth := 85.84

	if !reflect.DeepEqual(expectedSplitInMonth, splitInMonth) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSplitInMonth, splitInMonth)
	}
}

func TestGetSumOfInstalment_inEighteenMonth_forSmartphone(t *testing.T) {
	var service Service

	period := types.EighteenMonth
	category := types.Smartphone
	phone := "+992935598877"

	sum, splitInMonth, err := service.GetSumOfInstalment(category, 1000.0, phone, period)
	if err != nil {
		t.Error(err)
	}

	expectedSum := 1060.0

	if !reflect.DeepEqual(expectedSum, sum) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSum, sum)
	}

	expectedSplitInMonth := 58.89

	if !reflect.DeepEqual(expectedSplitInMonth, splitInMonth) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSplitInMonth, splitInMonth)
	}
}

func TestGetSumOfInstalment_inTwentyFourMonth_forSmartphone(t *testing.T) {
	var service Service

	period := types.TwentyFourMonth
	category := types.Smartphone
	phone := "+992935598877"

	sum, splitInMonth, err := service.GetSumOfInstalment(category, 1000.0, phone, period)
	if err != nil {
		t.Error(err)
	}

	expectedSum := 1090.0

	if !reflect.DeepEqual(expectedSum, sum) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSum, sum)
	}

	expectedSplitInMonth := 45.42

	if !reflect.DeepEqual(expectedSplitInMonth, splitInMonth) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSplitInMonth, splitInMonth)
	}
}

func TestGetSumOfInstalment_inEigteenMonth_forLaptop(t *testing.T) {
	var service Service

	period := types.EighteenMonth
	category := types.Laptop
	phone := "+992935598877"

	sum, splitInMonth, err := service.GetSumOfInstalment(category, 1000.0, phone, period)
	if err != nil {
		t.Error(err)
	}

	expectedSum := 1040.0

	if !reflect.DeepEqual(expectedSum, sum) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSum, sum)
	}

	expectedSplitInMonth := 57.78

	if !reflect.DeepEqual(expectedSplitInMonth, splitInMonth) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSplitInMonth, splitInMonth)
	}
}

func TestGetSumOfInstalment_inTwentyFourMonth_forLaptop(t *testing.T) {
	var service Service

	period := types.TwentyFourMonth
	category := types.Laptop
	phone := "+992935598877"
	// var accountID int64 = 1

	sum, splitInMonth, err := service.GetSumOfInstalment(category, 1000.0, phone, period)
	if err != nil {
		t.Error(err)
	}

	expectedSum := 1080.0

	if !reflect.DeepEqual(expectedSum, sum) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSum, sum)
	}

	expectedSplitInMonth := 45.0

	if !reflect.DeepEqual(expectedSplitInMonth, splitInMonth) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSplitInMonth, splitInMonth)
	}
}

func TestGetSumOfInstalment_inTwentyFourMonth_forTV(t *testing.T) {
	var service Service

	period := types.TwentyFourMonth
	category := types.TV
	phone := "+992935598877"

	sum, splitInMonth, err := service.GetSumOfInstalment(category, 1000.0, phone, period)
	if err != nil {
		t.Error(err)
	}

	expectedSum := 1050.0

	if !reflect.DeepEqual(expectedSum, sum) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSum, sum)
	}

	expectedSplitInMonth := 43.75

	if !reflect.DeepEqual(expectedSplitInMonth, splitInMonth) {
		t.Errorf("invalid result, expected: %v, actual: %v", expectedSplitInMonth, splitInMonth)
	}
}
