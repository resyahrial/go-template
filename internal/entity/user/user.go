package user

import "github.com/segmentio/ksuid"

type User struct {
	Ksuid ksuid.KSUID
	Name  string
}
