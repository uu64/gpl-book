// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func Pack(url string, ptr interface{}) (string, error) {
	v := reflect.ValueOf(ptr).Elem()
	fields := make(map[string][]string)

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		var key string
		key = tag.Get("http")
		if key == "" {
			key = strings.ToLower(fieldInfo.Name)
		}
		if _, ok := fields[key]; !ok {
			fields[key] = []string{}
		}

		v, err := depopulate(v.Field(i))
		if err != nil {
			return "", err
		}
		fields[key] = append(fields[key], v...)
	}

	query := []string{}
	for k, v := range fields {
		for _, q := range v {
			query = append(query, fmt.Sprintf("%s=%s", k, q))
		}
	}
	if len(query) > 0 {
		url = fmt.Sprintf("%s?%s", url, strings.Join(query, "&"))
	}
	return url, nil
}

func depopulate(v reflect.Value) ([]string, error) {
	values := []string{}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		values = append(values, fmt.Sprintf("%d", v.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		values = append(values, fmt.Sprintf("%d", v.Uint()))

	case reflect.Bool:
		values = append(values, fmt.Sprintf("%t", v.Bool()))

	case reflect.String:
		values = append(values, v.String())

	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			v, err := depopulate(v.Index(i))
			if err != nil {
				return nil, err
			}
			values = append(values, v...)
		}

	default:
		return nil, fmt.Errorf("unsupported kind %v", v.Kind())
	}
	return values, nil
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
