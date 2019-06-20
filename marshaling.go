package hl7

import (
	"errors"
	"reflect"
)

const tagName = "hl7"

func Unmarshal(m Message, out interface{}) error {
	t := reflect.TypeOf(out).Elem()
	v := reflect.ValueOf(out).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)

		result, ok, err := m.Query(tag)

		if err != nil {
			return err
		}

		if !ok {
			return errors.New("Could not get element")
		}

		v.Field(i).SetString(result)
	}

	return nil
}
