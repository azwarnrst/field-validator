package router

type UserFormData struct {
	UserName          string `form:"user_name" mandatory:"true"`
	UserID            int    `form:"user_id" mandatory:"true"`
	UserMail          string `form:"email" mandatory:"true"`
	UserFullName      string `form:"full_name"`
	UserAddress       string `form:"address"`
	IsOfficialPartner bool   `form:"is_partner" mandatory:"true"`
}
