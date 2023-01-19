package validators

import (
	"fmt"
)

type Validator struct {
	minParameterSize int
	maxParameterSize int
}

func New(min, max int) Validator {
	return Validator{min, max}
}

func (v Validator) ValidSizeParameter(a int) error {
	if a < v.minParameterSize || a > v.maxParameterSize {
		return fmt.Errorf("need to set parameters from %v to %v", v.minParameterSize, v.maxParameterSize)
	}

	return nil
}

func (v Validator) ValidateGameParameters(c, mines int) error {
	if mines < v.minParameterSize || mines > c {
		return fmt.Errorf("need to set count of mines from %v to %v", v.minParameterSize, c)
	}

	return nil
}

// GetMinParameters show parameters.
func (v Validator) GetMinParameters() int {
	return v.minParameterSize
}

// GetMaxParameters show parameters.
func (v Validator) GetMaxParameters() int {
	return v.maxParameterSize
}
