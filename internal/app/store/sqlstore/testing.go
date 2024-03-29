package sqlstore

import (
	"database/sql"
	"log"
	"testing"
)

func TestDB(t *testing.T, StorageInfo string) (*sql.DB, func(...string)) {
	t.Helper()
	db, err := sql.Open("postgres", StorageInfo)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE TABLE users")
			if err != nil {
				log.Fatal(err)
			}
		}

		db.Close()
	}
}
