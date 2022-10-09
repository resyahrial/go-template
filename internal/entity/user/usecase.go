package user

import "context"

type Usecase interface {
	GetAll(ctx context.Context) (users []User, err error)
}
