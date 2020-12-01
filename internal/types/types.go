package types

type UserFormData struct {
	//validate format : word = alphanumeric, alpha = alphabetic only, email = email format
	//min_length: minimum character length
	UserName          string `json:"user_name" name:"user_name" required:"true" format:"word" min_length:"6"` //these struct tags can be improved later with better tag grouping
	UserID            int    `json:"user_id" name:"user_id" required:"true"`
	UserMail          string `json:"email" name:"email" required:"true" format:"email" min_length:"5"`
	UserFullName      string `json:"full_name" name:"full_name" format:"alpha"`
	UserAddress       string `json:"address" name:"address" format:"word"`
	IsOfficialPartner bool   `json:"is_partner" name:"is_partner" required:"true"`
}

type HttpResp struct {
	Header HeaderResp
	Data   interface{} `json:"data"`
}

type HeaderResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
