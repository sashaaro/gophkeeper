package entity

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/sashaaro/gophkeeper/internal/utils"
)

type CreditCard struct {
	Number string
	Date   string
	Name   string
	Code   string
}

var (
	RegexpCardNumber = regexp.MustCompile(`^\d{16}$`)
	RegexpCardCode   = regexp.MustCompile(`^\d{3}$`)
	RegexpName       = regexp.MustCompile(`^[A-Z ]+$`)
	RegexpDate       = regexp.MustCompile(`^(0\d|1[012])/\d{2}$`)
)

func (c *CreditCard) Valid() error {
	var errs []error
	if err := c.validateNumber(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validateName(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validateDate(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validateCode(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (c *CreditCard) validateNumber() error {
	if !RegexpCardNumber.MatchString(c.Number) {
		return fmt.Errorf("card number should to contain 16 digits: %w", ErrValidation)
	}
	if !utils.CheckMoonAlgorithm(c.Number) {
		return fmt.Errorf("card number should to cover with Moon Alg: %w", ErrValidation)
	}
	return nil
}

func (c *CreditCard) validateDate() error {
	if !RegexpDate.MatchString(c.Date) {
		return fmt.Errorf("card date should have format MM/YY: %w", ErrValidation)
	}
	return nil
}

func (c *CreditCard) validateName() error {
	if !RegexpName.MatchString(c.Name) {
		return fmt.Errorf("card client name should contain only letters A-Z and spaces: %w", ErrValidation)
	}
	return nil
}

func (c *CreditCard) validateCode() error {
	if !RegexpCardCode.MatchString(c.Code) {
		return fmt.Errorf("card code should contaion only 3 digits: %w", ErrValidation)
	}
	return nil
}
