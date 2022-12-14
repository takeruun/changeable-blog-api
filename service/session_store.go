package service

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/wader/gormstore/v2"
)

type SessionStoreServiceRepository interface {
	GetSession(ctx context.Context, name string) (*sessions.Session, error)
	SaveSession(ctx context.Context, session *sessions.Session) error
	GetValue(ctx context.Context, key string) (interface{}, error)
	SaveValue(ctx context.Context, key string, value interface{}) error
}

type SessionStoreService struct {
	store *gormstore.Store
}

func NewSessionStoreService(s *gormstore.Store) SessionStoreServiceRepository {
	return &SessionStoreService{
		store: s,
	}
}

// HTTPKey is the key used to extract the Http struct.
type HTTPKey string

// HTTP is the struct used to inject the response writer and request http structs.
type HTTP struct {
	W *http.ResponseWriter
	R *http.Request
}

// GetSession returns a cached session of the given name
func (service *SessionStoreService) GetSession(ctx context.Context, name string) (*sessions.Session, error) {
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	// Ignore err because a session is always returned even if one doesn't exist
	session, err := service.store.Get(httpContext.R, name)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// SaveSession saves the session by writing it to the response
func (service *SessionStoreService) SaveSession(ctx context.Context, session *sessions.Session) error {
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	err := service.store.Save(httpContext.R, *httpContext.W, session)

	return err
}

func (service *SessionStoreService) GetValue(ctx context.Context, key string) (interface{}, error) {
	session, err := service.GetSession(ctx, "session")
	if err != nil {
		return "", err
	}

	return session.Values[key], nil
}

func (service *SessionStoreService) SaveValue(ctx context.Context, key string, value interface{}) error {
	session, err := service.GetSession(ctx, "session")
	if err != nil {
		return err
	}

	session.Values[key] = value

	err = service.SaveSession(ctx, session)
	if err != nil {
		return err
	}

	return nil
}
