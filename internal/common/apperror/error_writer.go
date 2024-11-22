package apperror

type ErrorWriter interface {
	WriteError(err error)
}
