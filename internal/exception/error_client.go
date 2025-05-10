package exception

type ErrorClientRequest struct {
	ActualError error
	StatusCode  int
	Detail      string
}

func NewErrorClientRequest(err error, statusCode int, detail string) *ErrorClientRequest {
	return &ErrorClientRequest{
		ActualError: err,
		StatusCode:  statusCode,
		Detail:      detail,
	}
}

func (e *ErrorClientRequest) Error() string {
	return e.ActualError.Error()
}

func (e *ErrorClientRequest) GetDetailError() string {
	return e.Detail
}

func (e *ErrorClientRequest) GetStatusCode() int {
	return e.StatusCode
}
