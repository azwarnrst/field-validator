package router

import (
	"encoding/json"
	"github.com/azwarnrst/field-validator/internal/types"
	"github.com/azwarnrst/field-validator/internal/validator"
	"log"
	"net/http"
	"strconv"
)

type XRouter struct {
	FormValidator validator.FormValidator
}

func (x *XRouter) UserHandler(w http.ResponseWriter, r *http.Request) {
	Uname := r.FormValue("user_name")
	UserID, err := strconv.Atoi(r.FormValue("user_id"))
	FullName := r.FormValue("full_name")
	Email := r.FormValue("email")
	Address := r.FormValue("address")
	IsPartner, err := strconv.ParseBool(r.FormValue("is_partner"))

	formData := types.UserFormData{
		UserName:          Uname,
		UserID:            UserID,
		UserMail:          Email,
		UserFullName:      FullName,
		UserAddress:       Address,
		IsOfficialPartner: IsPartner,
	}
	data, err := json.Marshal(&formData)
	if err != nil {
		log.Printf("error marshal form data, err %+v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (x *XRouter) UserHandler2(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		data        []byte
		respData    types.HttpResp
		valFormData = types.UserFormData{}
	)

	err = validator.NewValidator().ValidateFormData(r, &valFormData)
	if err != nil {
		log.Printf("invalid form data  : %+v", err)
		respData.Header.Status = http.StatusBadRequest
		respData.Header.Message = err.Error()
	} else {
		respData.Header.Status = http.StatusOK
		respData.Data = valFormData
	}

	data, err = json.Marshal(&respData)
	if err != nil {
		log.Printf("error marshal form data, err %+v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respData.Header.Status)
	w.Write(data)
}
