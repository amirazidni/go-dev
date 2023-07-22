package utils

import (
	"errors"

	"log"
)

var (
	ServerError         = GenerateError("Something went wrong! Please try again later")
	UserNotExist        = GenerateError("User not exists")
	UnauthorisedError   = GenerateError("You are not authorised to perform this action")
	TimeStampError      = GenerateError("time should be a unix timestamp")
	InternalServerError = GenerateError("internal server error")
)

func GenerateError(err string) error {
	return errors.New(err)
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
