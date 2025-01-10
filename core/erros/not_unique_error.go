package erros

import "fmt"

type NotUniqueError struct {
	tipo  string
	valor string
}

func NewNotUniqueError(tipo string, valor string) error {
	return NotUniqueError{
		tipo:  tipo,
		valor: valor,
	}
}

func (n NotUniqueError) Error() string {
	return fmt.Sprintf("O %s: %s já está sendo utilizado", n.tipo, n.valor)
}
