package exception

type CustomError struct {
	StatusCode int
	Err        error
}

func (r *CustomError) Error() string {
	return r.Err.Error()
}
