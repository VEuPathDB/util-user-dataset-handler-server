package command

func NewUserError(msg string) *UserError {
	return &UserError{msg}
}

type UserError struct {
	Message string
}

func (u *UserError) Error() string {
	return u.Message
}

func NewHandlerError(msg string) *HandlerError {
	return &HandlerError{msg}
}

type HandlerError struct {
	Message string
}

func (u *HandlerError) Error() string {
	return u.Message
}

