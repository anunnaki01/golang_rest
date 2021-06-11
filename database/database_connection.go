package database

import (
	"database/sql"
	"fmt"
	"os"
)

func GetDB() (*sql.DB, error) {
	return sql.Open(os.Getenv("DATABASE_DRIVER"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME")),
	)
}
