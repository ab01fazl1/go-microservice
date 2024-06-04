package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "attemped to create a record with an existing key"
}