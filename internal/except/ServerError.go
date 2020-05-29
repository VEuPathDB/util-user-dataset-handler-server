package except

func NewServerError(msg string) *ServerError {
	return &ServerError{msg}
}

type ServerError struct {
	Message string
}

func (s *ServerError) Error() string {
	return s.Message
}


