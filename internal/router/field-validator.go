package router

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type FormValidator struct {
	formData url.Values
}

func NewValidator() *FormValidator {
	return &FormValidator{}
}

func (v *FormValidator) validateType(label, data string, dataVal reflect.Value) (parseData reflect.Value, err error) {
	var (
		boolParse      bool
		intParse       int
		intDataParse   int64
		floatDataParse float64
	)

	defer func() {
		parseData = dataVal
		switch err {
		case strconv.ErrSyntax:
			err = errors.New("invalid numeric value for " + label)

		}

	}()

	log.Printf("kind : %+v\n", dataVal.Kind())

	switch dataVal.Kind() {
	case reflect.Bool:
		boolParse, err = strconv.ParseBool(data)
		dataVal.SetBool(boolParse)
	case reflect.Int:
		intParse, err = strconv.Atoi(data)
		if err == nil {
			intDataParse = int64(intParse)
			dataVal.SetInt(intDataParse)
		}
	case reflect.Int32:
		intDataParse, err = strconv.ParseInt(data, 10, 32)
		if err == nil {
			dataVal.SetInt(intDataParse)
		}

	case reflect.Float32:
		floatDataParse, err = strconv.ParseFloat(data, 32)
		if err == nil {
			dataVal.SetFloat(floatDataParse)
		}

	default:
		dataVal.SetString(data)
	}

	return
}

func (v *FormValidator) ValidateFormData(r *http.Request, data interface{}) (err error) {
	err = r.ParseForm()
	if err != nil {
		return errors.New("failed to parse form data")
	}

	if reflect.ValueOf(data).Kind() != reflect.Ptr {
		return errors.New("invalid data type")
	}

	if reflect.ValueOf(data).Elem().Kind() != reflect.Struct {
		return errors.New("invalid form struct")
	}

	val := reflect.Indirect(reflect.ValueOf(data))
	dataElement := reflect.ValueOf(data).Elem()
	for i := 0; i < val.NumField(); i++ {
		tagName := val.Type().Field(i).Tag.Get("name")
		if tagName == "" {
			err = errors.New("struct tag:name cannot be empty")
			break
		}

		formVal := r.FormValue(tagName)
		if strings.TrimSpace(formVal) == "" {
			if val.Type().Field(i).Tag.Get("required") == "true" {
				err = errors.New(tagName + " is mandatory")
			} else {
				continue
			}
		}

		if dataElement.Field(i).CanSet() {
			data, parseErr := v.validateType(tagName, formVal, val.Field(i))
			if parseErr != nil {
				err = parseErr
				break
			}
			dataElement.Field(i).Set(data)
		}
	}

	return

}
