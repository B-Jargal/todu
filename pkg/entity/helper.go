package entity

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

const (
	ROLE_OWNER = "owner"
	ROLE_ADMIN = "admin"
)

type Filter struct {
	Keyword string
	Role    string
}
