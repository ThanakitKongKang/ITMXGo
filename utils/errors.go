package util

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewCustomerNotFoundError() *NotFoundError {
	return &NotFoundError{Message: "customer not found"}
}
