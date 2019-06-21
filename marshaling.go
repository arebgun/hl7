package hl7

import (
	"errors"
	// "fmt"
	"reflect"
	"strconv"
	"time"
)

const tagName = "hl7"
const DefaultTimeFormat = "20060102150405.0000-0700"

var TimeFormat = DefaultTimeFormat

func Unmarshal(m Message, out interface{}) error {
	typeOf := reflect.TypeOf(out).Elem()
	valueOf := reflect.ValueOf(out).Elem()

	for i := 0; i < typeOf.NumField(); i++ {
		path, ok := typeOf.Field(i).Tag.Lookup(tagName)

		if !ok {
			if typeOf.Field(i).Type.Kind() == reflect.Struct {
				Unmarshal(m, valueOf.Field(i).Addr().Interface())
			}

			continue
		}

		result, ok, err := m.Query(path)

		if !ok {
			if err != nil {
				return err
			}

			return errors.New("Query result not ok")
		}

		switch typeOf.Field(i).Type.Kind() {
		case reflect.String:
			valueOf.Field(i).SetString(result)
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			value, err := strconv.ParseInt(result, 0, 64)
			if err != nil {
				return err
			}
			valueOf.Field(i).SetInt(value)
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
			value, err := strconv.ParseUint(result, 0, 64)
			if err != nil {
				return err
			}
			valueOf.Field(i).SetUint(value)
		case reflect.Float32:
		case reflect.Float64:
			value, err := strconv.ParseFloat(result, 64)
			if err != nil {
				return err
			}
			valueOf.Field(i).SetFloat(value)
		case reflect.Struct:
			if typeOf.Field(i).Type.PkgPath() == "time" && typeOf.Field(i).Type.Name() == "Time" {
				value, err := time.Parse(TimeFormat[:len(result)], result)
				if err != nil {
					return err
				}
				valueOf.Field(i).Set(reflect.ValueOf(value))
			} else {
				return errors.New("Could not handle " + typeOf.Field(i).Name)
			}
		default:
			return errors.New("Could not handle " + typeOf.Field(i).Name)
		}
	}

	return nil
}
