package models

type User struct {
	_id   string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`

	Password  string `"json:password,omitempty"`
	Title     string `"json:title,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	// FirstName string `json:"firstname,omitempty"`
	// LastName  string `json:"lastname,omitempty"`
	//Age       int    `json:"age,omitempty"`
	Roleid int    `json:"roleid,omitempty"`
	Token  string `"json:password,omitempty"`
}

type LoginDetails struct {
	Email    string
	Password string
}

type ErrorMessage struct {
	ErrorMessage string `json:"ErrorMessage"`
}
