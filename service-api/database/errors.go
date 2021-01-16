package database

// NotFoundError ...
type NotFoundError struct {
	Name string
}

func (err *NotFoundError) Error() string {
	return err.Name
}

//DuplicateError ...
type DuplicateError struct {
	Name string
}

func (err *DuplicateError) Error() string {
	return err.Name
}
