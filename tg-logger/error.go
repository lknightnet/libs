package tg_logger

import "errors"

func NewError(path string, err error) error {
	return errors.New(path + ": " + err.Error())
}
