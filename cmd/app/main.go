// A GORM demo. NOT A SERVICE: nothing listens on $PORT, so START_CMD is empty
// and bin/run stops at the start step by design.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Qode-Platform/qode-gorm-template-v1/store"
)

func main() {
	// Default to an in-memory database: the fleet's Go image is distroless with
	// a read-only /app, so a file DSN would fail at startup.
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "file::memory:"
	}
	db, err := store.Open(dsn)
	if err != nil {
		log.Fatal(err)
	}
	u, err := store.CreateUser(db, "Ada", fmt.Sprintf("ada+%d@example.com", 1))
	if err != nil {
		log.Println("create:", err)
	} else {
		log.Printf("created user %d (%s)", u.ID, u.Email)
	}
	var n int64
	db.Model(&store.User{}).Count(&n)
	log.Printf("users in table: %d", n)
}
