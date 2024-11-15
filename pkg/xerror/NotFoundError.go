package xerror

import "fmt"

type NotFoundError struct {
	Message string
	ID      int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf(e.Message, e.ID)
}
