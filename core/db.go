package core

import (
	"time"

	"github.com/boltdb/bolt"
)

// OpenDB open bolt db
func OpenDB() (*bolt.DB, error) {

	if DB == nil {
		db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			return new(bolt.DB), err
		}
		DB = db
	}

	return DB, nil
}
