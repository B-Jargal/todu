package application

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/B-Jargal/todu.git/common/oapi"
	"github.com/go-chi/chi"
)

func (app *Application) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var loggedUserID int
		switch {
		case app.Debug:
			loggedUserID = 1

		case !app.Debug:
			token := r.Header.Get("token")
			if token == "" {
				cookie, err := r.Cookie("token")
				if err != nil {
					oapi.ClientError(w, http.StatusUnauthorized)
					return
				}
				token = cookie.Value
			}

			loggedUserID = app.Blogin.Validate(token)
			if loggedUserID == 0 {
				oapi.ClientError(w, http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(r.Context(), ContextKeyIsAuthenticated, true)
		ctx = context.WithValue(ctx, ContextKeyAuthCustomer, customer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *Application) AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		clientID := r.Header.Get("client_id")
		clientSecret := r.Header.Get("client_secret")

		client, err := app.Clients.Authenticate(clientID, clientSecret)
		if err != nil {
			if errors.Is(err, clientman.ErrNotFound) {
				oapi.Forbidden(w)
			} else {
				oapi.ServerError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyIsAuthenticated, true)
		ctx = context.WithValue(ctx, ContextKeyAuthClient, client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *Application) AuthenticateClient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		key := r.Header.Get("key")

		if key != "6ltPXDZDCaoJRxsQbsGS84URkmZ4FTcefA70lCalPgh3e4g0QhGwbCOHV0hfcIF6zT2IGkWjle9tW7tVzrdlsnBHxtqFjavJhLaQkpao8Of7qkO9rC3LOQGrsDn7KD4S" {
			app.InfoLog.Println("Invalid key:", key)
			oapi.NotFound(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := r.Context().Value(ContextKeyAuthCustomer).(*customerman.Customer)
		if s.Role != customerman.ROLE_ADMIN {
			oapi.ClientError(w, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *Application) SetChosenOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggedCustomer := r.Context().Value(ContextKeyAuthCustomer).(*customerman.Customer)
		customerID, err := strconv.Atoi(chi.URLParam(r, "CustomerID"))
		if err != nil {
			oapi.NotFound(w)
			return
		}

		if loggedCustomer.Role != customerman.ROLE_ADMIN {
			if customerID != loggedCustomer.ID {
				oapi.ClientError(w, http.StatusForbidden)
				return
			}
		}

		customer, err := app.Customers.Get(customerID)
		if err != nil {
			if errors.Is(err, customerman.ErrNotFound) {
				oapi.NotFound(w)
			} else {
				oapi.ServerError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyChosenOne, customer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
