package exception

// AppError ...
type AppError struct {
	Description string
	Err error
}

func (err *AppError) Error() string {
	return err.Description
}
