package test_utils

import (
	"app/config"
	"errors"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(t *testing.T) *gorm.DB {
	c := config.NewConfig()
	db, err := gorm.Open(
		mysql.Open(c.DB.Username+":"+c.DB.Password+"@tcp("+c.DB.Host+")/"+c.DB.DBName+"?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{},
	)
	if err != nil {
		t.Fatal(err)
		panic("fail setup db")
	}

	return db
}

func RunInTransaction(db *gorm.DB, fn func(db *gorm.DB)) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		defer func() {
			tx.Rollback()
		}()
		fn(tx)
		return nil
	}); err == nil {
		return errors.New("Failed to Rollback")
	}
	return nil
}
