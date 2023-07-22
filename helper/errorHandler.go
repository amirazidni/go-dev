package helper

import (
	"fmt"
	"log"
	"net/http"
)

type ServerError struct {
	Code int
}

const (
	TimeStampError = iota
	BadRequestError
	UnauthorizedError
	ForbiddenError
	NotFoundError
	DatabaseError
	InternalServerError
)

func (e ServerError) Error() string {
	switch e.Code {
	case TimeStampError:
		return "time should be a unix timestamp"
	default:
		return "Unknown order error code"
	}
}

func (e ServerError) Unwrap() error {
	switch e.Code {
	case TimeStampError:
		return fmt.Errorf("%s: %v", e.Error(), http.StatusUnprocessableEntity)
	default:
		return fmt.Errorf("%s: %v", e.Error(), http.StatusInternalServerError)
	}
}
func ErrorHandler(err error, message string) bool {
	if err != nil {
		log.Printf(message + ". Cause: " + err.Error())
		return true
	}
	return false
}

func ErrorHandlerFatal(err error, message string) bool {
	if err != nil {
		log.Printf(message + ". Cause: " + err.Error())
		panic(err)
	}
	return false
}
