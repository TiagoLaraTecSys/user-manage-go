package erros

import (
	"fmt"
	"strings"
)

type InvalidRequestErr struct {
	campos []string
}

func NewInvalidRequestErr(campos ...string) error {
	return InvalidRequestErr{
		campos: campos,
	}
}

func (c InvalidRequestErr) Error() string {
	if len(c.campos) == 0 {
		return "requisição inválida!"
	} else {
		return fmt.Sprintf("requisição inválida!\n %s", strings.Join(c.campos, ""))
	}

}
