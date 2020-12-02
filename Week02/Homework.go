package main

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"os"
	"time"
	"fmt"
)

var (
	ctx context.Context
	db  *sql.DB
)
func DoDao(username string,created time.Time) (int, error) {


	err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users limit 1").Scan(&username, &created)


	if err != nil {
		return 0, errors.Wrap(err, "Dao No Data")
	}

	defer db.Close()
	return 1, nil
}
func Server() (int, error) {
	var username string
	var created time.Time
	count, err := DoDao(username,created)
	return count, errors.WithMessage(err,"Server No Data")
}
func main() {
	_, err := Server()
	if err != nil {
		fmt.Printf("original error: %T %v\n",errors.Cause(err),errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n",err)
		os.Exit(1)
	}
}
