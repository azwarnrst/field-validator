package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type XRouter struct {
	FormValidator FormValidator
}

func (x *XRouter) UserHandler(w http.ResponseWriter, r *http.Request) {
	Uname := r.FormValue("user_name")
	UserID, err := strconv.Atoi(r.FormValue("user_id"))
	FullName := r.FormValue("full_name")
	Email := r.FormValue("email")
	Address := r.FormValue("address")
	IsPartner, err := strconv.ParseBool(r.FormValue("is_partner"))

	formData := UserFormData{
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

	formData := UserFormData{}
	//err := json.NewDecoder(r.Body).Decode(&formData)
	//err := x.FormValidator.ValidateFormData(formData)
	//if err != nil {
	//
	//}

	Uname := r.FormValue("user_name")
	UserID, err := strconv.Atoi(r.FormValue("user_id"))
	FullName := r.FormValue("full_name")
	Email := r.FormValue("email")
	Address := r.FormValue("address")
	IsPartner, err := strconv.ParseBool(r.FormValue("is_partner"))

	//log.Printf("%+v", r.Form)

	formData := UserFormData{
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
