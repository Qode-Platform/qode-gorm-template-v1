// Package store holds the models and the data access built on GORM.
//
// The driver is glebarez/sqlite, a PURE GO sqlite. That matters: the fleet's
// Go image builds with CGO_ENABLED=0 into distroless, and the cgo-based
// mattn/go-sqlite3 cannot link there.
package store

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:120;not null"`
	Email string `gorm:"size:255;uniqueIndex;not null"`
}

// Open opens the database and applies the schema. dsn "file::memory:" gives a
// throwaway database, which is what the tests use.
func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateUser(db *gorm.DB, name, email string) (*User, error) {
	u := &User{Name: name, Email: email}
	return u, db.Create(u).Error
}

func FindByEmail(db *gorm.DB, email string) (*User, error) {
	var u User
	return &u, db.Where("email = ?", email).First(&u).Error
}
