package hl7

import (
	"errors"
	_ "fmt"
	"reflect"
	"strconv"
	"time"
)

func unpack(out interface{}, query func(q *Query) (string, bool), querySlice func(q *Query) []string) (err error) {
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
				unpack(valueOf.Field(i).Addr().Interface(), query, querySlice)
			}

			continue
		}

		q, err := ParseQuery(path)
		panicOnError(err)

		switch typeOf.Field(i).Type.Kind() {
		case reflect.String:
			result, _ := query(q)
			valueOf.Field(i).SetString(result)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result, _ := query(q)
			value, err := strconv.ParseInt(result, 0, 64)
			panicOnError(err)

			valueOf.Field(i).SetInt(value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result, _ := query(q)
			value, err := strconv.ParseUint(result, 0, 64)
			panicOnError(err)

			valueOf.Field(i).SetUint(value)
		case reflect.Float32, reflect.Float64:
			result, _ := query(q)
			value, err := strconv.ParseFloat(result, 64)
			panicOnError(err)

			valueOf.Field(i).SetFloat(value)
		case reflect.Slice, reflect.Array:
			switch typeOf.Field(i).Type.Elem().Kind() {
			case reflect.String:
				result := querySlice(q)
				valueOf.Field(i).Set(reflect.ValueOf(result))
			default:
				return errors.New("Could not handle " + typeOf.Field(i).Name)
			}
		case reflect.Struct:
			result, _ := query(q)
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

func (m Message) Unpack(out interface{}) (err error) {
	query := func(q *Query) (string, bool) {
		return m.query(q), true
	}

	querySlice := func(q *Query) []string {
		return m.querySlice(q)
	}

	return unpack(out, query, querySlice)
}

func (s Segment) Unpack(out interface{}) (err error) {
	query := func(q *Query) (string, bool) {
		return s.query(q), true
	}

	querySlice := func(q *Query) []string {
		return s.querySlice(q)
	}

	return unpack(out, query, querySlice)
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
