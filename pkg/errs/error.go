package errs

import (
	"bytes"
	"fmt"
)

type BasicError struct {
	Code           string        `json:"code"`
	Msg            string        `json:"msg"`
	HTTPStatusCode int           `json:"-"`
	Vs             []interface{} `json:"-"`
}

func (e *BasicError) clone() *BasicError {
	return &BasicError{
		Code:           e.Code,
		Msg:            e.Msg,
		HTTPStatusCode: e.HTTPStatusCode,
		Vs:             e.Vs,
	}
}

func (e *BasicError) Error() string {
	return e.Msg
}

func (e *BasicError) Message(msg string) *BasicError {
	err := e.clone()
	err.Msg = msg
	return err
}

func (e *BasicError) HTTPCode(code int) *BasicError {
	err := e.clone()
	err.HTTPStatusCode = code
	return err
}

func (e *BasicError) Params(params ...interface{}) *BasicError {
	err := e.clone()
	err.Vs = params
	return err
}

func (e *BasicError) Parse() {
	e.Msg = fmt.Sprintf(e.Msg, e.Vs...)
}

type Prefix string

func (p Prefix) Code(code string) *BasicError {
	var buffer bytes.Buffer
	buffer.WriteString(string(p))
	buffer.WriteString("-")
	buffer.WriteString(code)
	err := &BasicError{}
	err.Code = buffer.String()
	return err
}
