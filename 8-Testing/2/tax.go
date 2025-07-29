package tax

import (
	"errors"
	"time"
)

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	return repository.SaveTax(tax)
}

func CalculateTax2(amount float64) (float64, error) {
	time.Sleep(time.Millisecond)
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}
	if amount >= 1000 {
		return 10.0, nil
	}
	return 5.0, nil
}

func CalculateTax(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount <= 0 {
		return 0.0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}