package sqlite

import (
	"database/sql/driver"

	"github.com/gokutils/uuid"
	"modernc.org/sqlite"
)

func init() {
	sqlite.MustRegisterDeterministicScalarFunction(
		"uuid_generate_v4",
		0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return uuid.New().Value()
		},
	)
	sqlite.MustRegisterDeterministicScalarFunction(
		"uuid_generate_v5",
		2,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return uuid.NewSha1(uuid.ParseOrNil(args[0].(string)), []byte(args[1].(string))).Value()
		},
	)
}
