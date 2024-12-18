package erros

import "fmt"

type ChangeIdadeErr struct {
	idade int
}

func NewChangeIdadeErr(idade int) error {
	return ChangeIdadeErr{idade: idade}
}

func (c ChangeIdadeErr) Error() string {
	return fmt.Sprintf("mudança de idade inválida! %v", c.idade)
}
