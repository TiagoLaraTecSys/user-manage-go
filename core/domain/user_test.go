package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	email := "laratecsys@gmail.com"
	idade := 26

	user, _ := NewUser(
		WithEmail(email),
		WithIdade(idade),
		WithName("Tiago"),
	)

	assert.Equal(t, user.Idade, 26)
}
