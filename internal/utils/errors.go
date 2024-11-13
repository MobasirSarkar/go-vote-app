package utils

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ERROR_SAME_EMAIL          = errors.New(`duplicate key value violates unique constraint "users_email_key"`)
	ERROR_NO_REPONSE          = errors.New("Response Is Empty")
	ERROR_INVALID_CREDENTIALS = errors.New("Invalid Credentials.Please provide a correct credentials.")
)

const UniqueViolationCode = "23505"

type DuplicateEntryError struct {
	Field string
	Value string
}

func (e *DuplicateEntryError) Error() string {
	return fmt.Sprintf("%s with value '%s' already exists", e.Field, e.Value)
}

func IsDuplicateEntryError(err error, dst *DuplicateEntryError) bool {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) && pgError.Code == UniqueViolationCode {
		*dst = DuplicateEntryError{Field: pgError.ConstraintName, Value: pgError.Message}
		return true
	}
	return false
}
