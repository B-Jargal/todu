package application

import (
	"log"
	"time"

	"github.com/B-Jargal/todu.git/pkg/entity"
)

type contextKey string

const (
	ContextKeyIsAuthenticated = contextKey("isAuthenticated")
	ContextKeyAuthCustomer    = contextKey("authenticatedCustomer")
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Location *time.Location
	Debug    bool
	DataBase *entity.Database
}
