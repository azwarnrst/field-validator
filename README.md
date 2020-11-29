## form field validation helper

simple form input validation helper, to simplify parse input form & validate data 


* example request 
```
curl --location --request POST 'localhost:8080' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'user_name=jaret' \
--data-urlencode 'user_id=10' \
--data-urlencode 'email=jaret@jaret.com' \
--data-urlencode 'address=semarang' \
--data-urlencode 'is_partner=false' \
--data-urlencode 'full_name=jaret kuproy'
```  

* use regular way to get formValue and put into struct
```

    Uname := r.FormValue("user_name")
	UserID, err := strconv.Atoi(r.FormValue("user_id"))
	FullName := r.FormValue("full_name")
	Email := r.FormValue("email")
	Address := r.FormValue("address")
	IsPartner, err := strconv.ParseBool(r.FormValue("is_partner"))
    .....
    // validate formValue ....
    .....
	userFormPayload := types.UserFormData{
		UserName:          Uname,
		UserID:            UserID,
		UserMail:          Email,
		UserFullName:      FullName,
		UserAddress:       Address,
		IsOfficialPartner: IsPartner,
	}

``` 

* use field validator helper to parse & validate form data and put into struct
```
type UserFormData struct {
	UserName          string `name:"user_name" required:"true"`
	UserID            int    `name:"user_id" required:"true"`
	UserMail          string `name:"email" required:"true"`
	UserFullName      string `name:"full_name"`
	UserAddress       string `name:"address"`
	IsOfficialPartner bool   `name:"is_partner" required:"true"`
}

userFormPayload := UserFormData{}
err := validator.NewValidator().ValidateFormData(r, &userFormPayload)

```

* struct tag notes
```
name        :  r.FormValue({name}), field name to get form value
required    : set to true if field is mandatory   
```

