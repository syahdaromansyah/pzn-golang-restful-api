package exception

type ErrorInternalServer struct {
	ActualError error
	Detail      string
}

func NewErrorInternalServer(actualError error, detail string) *ErrorInternalServer {
	return &ErrorInternalServer{
		ActualError: actualError,
		Detail:      detail,
	}
}

func (e *ErrorInternalServer) DetailError() string {
	return e.Detail
}

func (e *ErrorInternalServer) Error() string {
	return e.ActualError.Error()
}
