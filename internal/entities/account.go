package entities

import "context"

type AccountType int

const (
	Regular AccountType = iota
	Premium
)

type Account struct {
	Id   string
	Type AccountType
}

type AccountUsecase interface {
	GetAccountById(ctx context.Context, id string) (account *Account, err error)
}
