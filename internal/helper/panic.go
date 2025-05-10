package helper

import "github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ClientPanicIfError(err, errorClient error) {
	if err != nil {
		panic(errorClient)
	}
}

func InternalServerPanicIfError(err error, detail string) {
	if err != nil {
		panic(exception.NewErrorInternalServer(err, detail))
	}
}
