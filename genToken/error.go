package genToken

import "errors"

func NewError(path string, err error) error {
	return errors.New(path + ": " + err.Error())
}
