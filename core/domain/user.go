package domain

import (
	"projeto-final/core/erros"
)

type (
	User struct {
		Id    int
		Email string
		Name  string
		Idade int
	}
	Opt func(*User)
)

func NewUser(opts ...Opt) (*User, error) {
	u := &User{}

	for _, opt := range opts {
		opt(u)
	}

	err := u.validarIdade()

	return u, err
}

func WithName(name string) Opt {
	return func(u *User) {
		u.Name = name
	}
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
