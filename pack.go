package hl7

import (
	"reflect"
	"strconv"
	"time"
)

func pack(in interface{}, setString func(query string, value string) error) error {
	typeOf := reflect.TypeOf(in).Elem()
	valueOf := reflect.ValueOf(in).Elem()

	for i := 0; i < typeOf.NumField(); i++ {
		path, ok := typeOf.Field(i).Tag.Lookup("hl7")

		if !ok {
			if typeOf.Field(i).Type.Kind() == reflect.Struct {
				pack(valueOf.Field(i).Addr().Interface(), setString)
			}

			continue
		}

		if _, err := ParseQuery(path); err != nil {
			return err
		}

		switch typeOf.Field(i).Type.Kind() {
		case reflect.String:
			str := valueOf.Field(i).String()

			if err := setString(path, str); err != nil {
				return err
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str := strconv.Itoa(int(valueOf.Field(i).Int()))

			if err := setString(path, str); err != nil {
				return err
			}
		// case reflect.Slice:
		// 	switch typeOf.Field(i).Type.Elem().Kind() {
		// 	case reflect.String:
		// 		valueOf.Field(i).Set(reflect.ValueOf(querySlice(q)))
		// 	}
		case reflect.Struct:
			if typeOf.Field(i).Type.PkgPath() == "time" && typeOf.Field(i).Type.Name() == "Time" {
				str := valueOf.Field(i).Elem().Interface().(time.Time).Format(timeFormat)

				if err := setString(path, str); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (m *Message) Pack(in interface{}) (err error) {
	setString := func(query string, value string) error {
		return m.SetString(query, value)
	}

	return pack(in, setString)
}
