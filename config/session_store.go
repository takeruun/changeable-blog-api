package config

import (
	"net/http"

	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

type SESSION_STORE struct {
	SECRET_HASH_KEY string
}

func NewSessionStore(db *gorm.DB) *gormstore.Store {
	c := NewConfig()
	return newSessionStore(db, &SESSION_STORE{
		SECRET_HASH_KEY: c.SESSION_STORE.Production.SecretHashKey,
	})
}

func newSessionStore(db *gorm.DB, ss *SESSION_STORE) *gormstore.Store {
	store := gormstore.NewOptions(
		db,
		gormstore.Options{},
		[]byte(ss.SECRET_HASH_KEY),
	)

	store.SessionOpts.Secure = true
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.MaxAge = 60 * 60 * 24 * 60
	store.SessionOpts.SameSite = http.SameSiteNoneMode

	return store
}
