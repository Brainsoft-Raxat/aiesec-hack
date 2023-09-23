package apperror

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	
	pkgerr "github.com/pkg/errors"
)

type ErrorMessage struct {
	Code    int
	Message string
}

type ErrorInfo struct {
	Status           int    `json:"status"`
	Code             int    `json:"-"`
	FullCode         string `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	error
}

func NewErrorInfo(ctx context.Context, errorCode ErrorCode, developerMessage string) *ErrorInfo {
	appErr := &ErrorInfo{
		Status:           errorCode.status,
		Code:             errorCode.code,
		Message:          errorCode.message,
		DeveloperMessage: developerMessage,
	}
	appErr.FullCode = appErr.fullCode()

	return appErr
}

func (e *ErrorInfo) fullCode() string {
	return strconv.Itoa(e.Status*100 + e.Code)
}

func (e *ErrorInfo) Error() string {
	if e.error == nil {
		return fmt.Sprintf("%s %s", e.fullCode(), e.DeveloperMessage)
	}

	return fmt.Sprintf("%s %s: %v", e.fullCode(), e.DeveloperMessage, e.error)
}

func (e *ErrorInfo) Cause() error {
	return e.error
}

// Unwrap для возможности errors.Is() и errors.As()
func (e *ErrorInfo) Unwrap() error {
	return e.error
}

// Wrap обертка ошибки
func (e *ErrorInfo) Wrap(err error) *ErrorInfo {
	if e.error != nil {
		err = pkgerr.Wrap(err, e.error.Error())
	}

	appErr := &ErrorInfo{
		Status:           e.Status,
		Code:             e.Code,
		Message:          e.Message,
		DeveloperMessage: e.DeveloperMessage,
		error:            err,
	}
	appErr.FullCode = appErr.fullCode()

	return appErr
}

func (e *ErrorInfo) copy() *ErrorInfo {
	err := *e

	return &err
}

func AsErrorInfo(err error) *ErrorInfo {
	var target *ErrorInfo
	if errors.As(err, &target) {
		return target
	}

	return nil
}


type ErrorCode struct {
	code    int
	status  int
	message string
}

func NewErrorCode(code, status int, message string) (response ErrorCode) {
	response.code = code
	response.status = status
	response.message = message

	return
}