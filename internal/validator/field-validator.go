package validator

import (
	"errors"
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
		uintDataParse  uint64
		floatDataParse float64
		intBitMap      = map[reflect.Kind]int{
			reflect.Int8:  8,
			reflect.Int16: 16,
			reflect.Int32: 32,
			reflect.Int64: 64,
		}
		uintBitMap = map[reflect.Kind]int{
			reflect.Uint8:  8,
			reflect.Uint16: 16,
			reflect.Uint32: 32,
			reflect.Uint64: 64,
		}

		floatBitMap = map[reflect.Kind]int{
			reflect.Float32: 32,
			reflect.Float64: 64,
		}
	)

	defer func() {
		parseData = dataVal
		if err != nil {
			err = errors.New("invalid data type for input " + label)
		}
	}()

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
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intDataParse, err = strconv.ParseInt(data, 10, intBitMap[dataVal.Kind()])
		if err == nil {
			dataVal.SetInt(intDataParse)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintDataParse, err = strconv.ParseUint(data, 10, uintBitMap[dataVal.Kind()])
		if err == nil {
			dataVal.SetUint(uintDataParse)
		}

	case reflect.Float32, reflect.Float64:
		floatDataParse, err = strconv.ParseFloat(data, floatBitMap[dataVal.Kind()])
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

	if reflect.ValueOf(data).Kind() != reflect.Ptr || reflect.ValueOf(data).Elem().Kind() != reflect.Struct {
		return errors.New("invalid struct form data")
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
				break
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
