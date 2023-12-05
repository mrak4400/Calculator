package calculator

import (
	"calculator/roman"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	Sum            = "+"
	Multiplication = "*"
	Division       = "/"
	Subtraction    = "-"
)

type Calculator struct {
	Input         string
	MathOperation string
	FirstValue    int
	SecondValue   int
}

func (c *Calculator) ParseInput() error {
	parts := strings.Fields(c.Input)
	if len(parts) != 3 {
		return errors.New("неверный формат ввода")
	}

	firstValueStr := parts[0]
	c.MathOperation = parts[1]
	secondValueStr := parts[2]
	var err error
	var isFirstRoman bool
	var isSecondRoman bool

	c.FirstValue, isFirstRoman = roman.RomanToArabic(firstValueStr)
	c.SecondValue, isSecondRoman = roman.RomanToArabic(secondValueStr)

	if isFirstRoman == true && isSecondRoman == true {
		return nil
	}

	if isFirstRoman != isSecondRoman {
		return errors.New("неправильный формат ввода")
	}

	c.FirstValue, err = strconv.Atoi(firstValueStr)
	if err != nil {
		return errors.New("ошибка при преобразовании первого значения")
	}

	c.SecondValue, err = strconv.Atoi(secondValueStr)
	if err != nil {
		return errors.New("ошибка при преобразовании второго значения")
	}

	return nil

}
func NewCalculator(input string) (*Calculator, error) {
	calculator := &Calculator{Input: input}
	err := calculator.ParseInput()
	if err != nil {
		return nil, err
	}

	return calculator, nil
}

func (c *Calculator) Calculate() int {
	switch c.MathOperation {
	case Sum:
		return c.FirstValue + c.SecondValue
	case Subtraction:
		return c.FirstValue - c.SecondValue
	case Multiplication:
		return c.FirstValue * c.SecondValue
	case Division:
		if c.SecondValue == 0 {
			fmt.Println("Ошибка деление на ноль")
			return 0
		}
		return c.FirstValue / c.SecondValue
	default:
		fmt.Println("Неверная математическая операция")
		return -1
	}
}
