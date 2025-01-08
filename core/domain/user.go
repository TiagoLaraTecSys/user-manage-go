package domain

import (
	"projeto-final/core/erros"
)

type (
	User struct {
		Id    string
		Email string
		Idade int
	}
	Opt func(*User)
)

func NewUser(opts ...Opt) *User {
	u := &User{}

	for _, opt := range opts {
		opt(u)
	}

	u.validarIdade()

	return u
}

func WithEmail(email string) Opt {

	return func(u *User) {
		u.Email = email
	}
}

func WithIdade(idade int) Opt {

	return func(u *User) {
		u.Idade = idade
	}
}

func (c *User) validarIdade() error {

	if c.Idade < 0 {
		return erros.NewChangeIdadeErr(c.Idade)
	}
	return nil
}
