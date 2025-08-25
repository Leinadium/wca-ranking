package repository

import (
	"database/sql"

	"github.com/guregu/null"
)

func SQLNullInt32(x sql.NullInt32) null.Int {
	return null.NewInt(int64(x.Int32), x.Valid)
}

func SQLNullString(x sql.NullString) null.String {
	return null.NewString(x.String, x.Valid)
}
