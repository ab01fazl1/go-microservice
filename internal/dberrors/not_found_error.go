package dberrors

import "fmt"

type NotFoundErro struct{
	Entity string
	ID string
}

func (e *NotFoundErro) Error() string {
	return fmt.Sprintf("unable to find %s with id %s", e.Entity, e.ID)
}