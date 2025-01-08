package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	email := "laratecsys@gmail.com"
	idade := 26

	user := NewUser(
		WithEmail(email),
		WithIdade(idade),
	)

	assert.Equal(t, user.Idade, 26)
}
