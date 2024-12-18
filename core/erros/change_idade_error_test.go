package erros

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeIdadeErr(t *testing.T) {
	test := struct {
		name  string
		idade int
		err   string
	}{
		name:  "Testando funcao de erro",
		idade: -10,
		err:   fmt.Sprintf("mudança de idade inválida! %v", -10),
	}
	t.Run(test.name, func(t *testing.T) {
		err := NewChangeIdadeErr(test.idade)
		assert.Equal(t, test.err, err.Error())
	})
}
