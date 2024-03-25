// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler).
// This file is meant to be re-generated in place and/or deleted at any time.

package order_dao

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.

type M map[string]interface{}

type whereHelpernull_String struct{ field string }
type whereHelperint struct{ field string }
type whereHelpertime_Time struct{ field string }
type whereHelperstring struct{ field string }
type whereHelpernull_Int struct{ field string }
type whereHelperint64 struct{ field string }

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("order_dao: failed to synchronize data after insert")

type insertCache struct {
	query        string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}