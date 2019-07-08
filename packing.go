package hl7

import (
	"errors"
	// "fmt"
	"reflect"
	"strconv"
	"time"
)

func (m Message) Unpack(out interface{}) (err error) {
	defer func() {
		if unknown := recover(); unknown != nil {
			err = unknown.(error)
		}
	}()

	typeOf := reflect.TypeOf(out).Elem()
	valueOf := reflect.ValueOf(out).Elem()

	for i := 0; i < typeOf.NumField(); i++ {
		path, ok := typeOf.Field(i).Tag.Lookup(tagName)

		if !ok {
			if typeOf.Field(i).Type.Kind() == reflect.Struct {
				m.Unpack(valueOf.Field(i).Addr().Interface())
			}

			continue
		}

		result, ok, err := m.Query(path)

		if !ok {
			panicOnError(err)
		}

		switch typeOf.Field(i).Type.Kind() {
		case reflect.String:
			valueOf.Field(i).SetString(result)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value, err := strconv.ParseInt(result, 0, 64)
			panicOnError(err)

			valueOf.Field(i).SetInt(value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value, err := strconv.ParseUint(result, 0, 64)
			panicOnError(err)

			valueOf.Field(i).SetUint(value)
		case reflect.Float32, reflect.Float64:
			value, err := strconv.ParseFloat(result, 64)
			panicOnError(err)

			valueOf.Field(i).SetFloat(value)
		case reflect.Struct:
			if typeOf.Field(i).Type.PkgPath() == "time" && typeOf.Field(i).Type.Name() == "Time" {
				if len(result) > 0 {
					value, err := time.ParseInLocation(timeFormat[:len(result)], result, Locale)
					panicOnError(err)

					valueOf.Field(i).Set(reflect.ValueOf(value))
				}
			} else {
				return errors.New("Could not handle " + typeOf.Field(i).Name)
			}
		default:
			return errors.New("Could not handle " + typeOf.Field(i).Name)
		}
	}

	return nil
}

const tagName = "hl7"
const timeFormat = "20060102150405.0000-0700"

var Locale *time.Location

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	Locale, _ = time.LoadLocation("")
}
