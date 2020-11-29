package types

type UserFormData struct {
	UserName          string `json:"user_name" name:"user_name" required:"true"`
	UserID            int    `json:"user_id" name:"user_id" required:"true"`
	UserMail          string `json:"email" name:"email" required:"true"`
	UserFullName      string `json:"full_name" name:"full_name"`
	UserAddress       string `json:"address" name:"address"`
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
