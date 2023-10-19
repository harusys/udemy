package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

// http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

var ErrRecordNotFound = pgx.ErrNoRows

var ErrForeignKeyViolation = &pq.Error{
	Code: ForeignKeyViolation,
}

var ErrUniqueViolation = &pq.Error{
	Code: UniqueViolation,
}

func ErrorCode(err error) string {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code.Name()
	}
	return ""
}
