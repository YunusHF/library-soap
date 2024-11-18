package app

import (
	"soap-library/pkg/pkguid"

	"github.com/hashicorp/go-multierror"
)

func (app *App) initSnowflakeGen() {
	snowflake, err := pkguid.NewSnowflake()
	if err != nil {
		app.err = multierror.Append(app.err, err)
	}

	app.snowflake = snowflake
}
