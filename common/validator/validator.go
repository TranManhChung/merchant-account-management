package validator

import (
	"fmt"
	"main.go/common/err"
)

type Validator struct {
	lens map[string]int
}

func New(lens map[string]int) Validator {
	return Validator{
		lens: lens,
	}
}

func (c Validator) CheckLen(fields map[string]int) error {
	for k, v := range fields {
		if v > c.lens[k] {
			return err.InternalError{Code: err.InvalidParameter.Code, Mess: fmt.Sprintf("%s is too long", k)}
		}
	}
	return nil
}

func (c Validator) CheckEmpty(fields map[string]interface{}) error {
	for k, v := range fields {
		switch val := v.(type) {
		case int64:
			if val == 0 {
				return err.InternalError{Code: err.InvalidParameter.Code, Mess: fmt.Sprintf("%s is empty", k)}
			}
		case string:
			if val == "" {
				return err.InternalError{Code: err.InvalidParameter.Code, Mess: fmt.Sprintf("%s is empty", k)}
			}
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	return nil
}
