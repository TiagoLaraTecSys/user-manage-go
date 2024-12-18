package domain

import (
	"projeto-final/core/erros"
)

type (
	User struct {
		Id    string
		email string
		idade int
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
		u.email = email
	}
}

func WithIdade(idade int) Opt {

	return func(u *User) {
		u.idade = idade
	}
}

func (c *User) validarIdade() error {

	if c.idade < 0 {
		return erros.NewChangeIdadeErr(c.idade)
	}
	return nil
}
