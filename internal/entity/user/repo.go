package user

import "context"

type Repo interface {
	GetAll(ctx context.Context) (users []User, err error)
}
