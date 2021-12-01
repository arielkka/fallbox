package errors

var (
	notFound = newError("not found")
)

type Errors struct {
	Description string
}

func newError(description string) *Errors {
	return &Errors{Description: description}
}

func (e Errors) Error() string {
	return e.Description
}

func NotFound() error {
	return notFound
}
