package service

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/wader/gormstore/v2"
)

// HTTPKey is the key used to extract the Http struct.
type HTTPKey string

// HTTP is the struct used to inject the response writer and request http structs.
type HTTP struct {
	W *http.ResponseWriter
	R *http.Request
}

// GetSession returns a cached session of the given name
func GetSession(ctx context.Context, store *gormstore.Store, name string) (*sessions.Session, error) {
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	// Ignore err because a session is always returned even if one doesn't exist
	session, err := store.Get(httpContext.R, name)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// SaveSession saves the session by writing it to the response
func SaveSession(ctx context.Context, store *gormstore.Store, session *sessions.Session) error {
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	err := store.Save(httpContext.R, *httpContext.W, session)

	return err
}
