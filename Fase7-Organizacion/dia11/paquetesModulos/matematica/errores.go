package matematica

import "errors"

var ErrDivisionCero = errors.New("no se puede dividir por cero")

func ErrorDivisionCero() error {
	return ErrDivisionCero
}
