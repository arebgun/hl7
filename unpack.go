package hl7

import (
	"reflect"
	"strconv"
	"time"
)

func unpack(out interface{}, query func(q *Query) string, querySlice func(q *Query) []string) error {
	typeOf := reflect.TypeOf(out).Elem()
	valueOf := reflect.ValueOf(out).Elem()

	for i := 0; i < typeOf.NumField(); i++ {
		path, ok := typeOf.Field(i).Tag.Lookup("hl7")

		if !ok {
			if typeOf.Field(i).Type.Kind() == reflect.Struct {
				unpack(valueOf.Field(i).Addr().Interface(), query, querySlice)
			}

			continue
		}

		q, err := ParseQuery(path)

		if err != nil {
			return err
		}

		switch typeOf.Field(i).Type.Kind() {
		case reflect.String:
			valueOf.Field(i).SetString(query(q))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value, err := strconv.ParseInt(query(q), 0, 64)

			if err != nil {
				return err
			}

			valueOf.Field(i).SetInt(value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value, err := strconv.ParseUint(query(q), 0, 64)

			if err != nil {
				return err
			}

			valueOf.Field(i).SetUint(value)
		case reflect.Float32, reflect.Float64:
			value, err := strconv.ParseFloat(query(q), 64)

			if err != nil {
				return err
			}

			valueOf.Field(i).SetFloat(value)
		case reflect.Slice:
			switch typeOf.Field(i).Type.Elem().Kind() {
			case reflect.String:
				valueOf.Field(i).Set(reflect.ValueOf(querySlice(q)))
			}
		case reflect.Struct:
			if typeOf.Field(i).Type.PkgPath() == "time" && typeOf.Field(i).Type.Name() == "Time" {
				result := query(q)

				if len(result) > 0 {
					value, err := time.ParseInLocation(timeFormat[:len(result)], result, Locale)

					if err != nil {
						return err
					}

					valueOf.Field(i).Set(reflect.ValueOf(value))
				}
			}
		}
	}

	return nil
}

func (m Message) Unpack(out interface{}) (err error) {
	query := func(q *Query) string {
		str, _ := m.query(q)
		return str
	}

	querySlice := func(q *Query) []string {
		return m.querySlice(q)
	}

	return unpack(out, query, querySlice)
}

func (s Segment) Unpack(out interface{}) (err error) {
	query := func(q *Query) string {
		str, _ := s.query(q)
		return str
	}

	querySlice := func(q *Query) []string {
		return s.querySlice(q)
	}

	return unpack(out, query, querySlice)
}

const timeFormat = "20060102150405.0000-0700"

var Locale *time.Location

func init() {
	Locale, _ = time.LoadLocation("")
}
