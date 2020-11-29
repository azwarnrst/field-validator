package router

import (
	"net/url"
)

type FormValidator struct {

}

func NewValidator(data  url.Values) *FormValidator {
	return &FormValidator{}
}

func (v *FormValidator) ValidateFormData(formFata interface{})(err error){

	return
}
