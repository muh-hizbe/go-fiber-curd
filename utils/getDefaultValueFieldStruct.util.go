package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func GetDefault(field string, file interface{}) (string, error) {
	typ := reflect.TypeOf(file)
	f, _ := typ.FieldByName(field)
	result := f.Tag.Get("default")

	if result == "" {
		return result, errors.New(fmt.Sprintf("Field %v not Found.", field))
	}

	return result, nil
}
