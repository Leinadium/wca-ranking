package utils

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

func NullStringtoPointer(x null.String) *string {
	return x.Ptr()
}

func NullIntToPointer(x null.Int) *int {
	var res *int = nil
	if x.Valid {
		res = new(int)
		*res = int(x.Int64)
	}
	return res
}

func NullResultsToSlice(x [5]null.Int) *[]int {
	res := []int{
		int(x[0].ValueOrZero()),
		int(x[1].ValueOrZero()),
		int(x[2].ValueOrZero()),
		int(x[3].ValueOrZero()),
		int(x[4].ValueOrZero()),
	}
	return &res
}
