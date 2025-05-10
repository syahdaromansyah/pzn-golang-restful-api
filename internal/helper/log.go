package helper

import (
	"log"
)

func LogStdPanicIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
