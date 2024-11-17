//nolint:unused // for later use
package app

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (app *App) initDB() error {
	connectionString := "root:AdminRoot@tcp(localhost:3306)/library?parseTime=true"

	var err error

	// Open database connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	app.closersFn = append(app.closersFn, func(_ context.Context) error {
		db.Close()
		return nil
	})

	if err := db.Ping(); err != nil {
		return err
	}

	app.database = db

	return nil
}
