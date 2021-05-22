package models

type User struct {
	_id       string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
	Title     string `"json:title,omitempty"`
	Roleid    int    `json:"roleid,omitempty"`
	Password  string `"json:password,omitempty"`
	Token     string `"json:password,omitempty"`
}

type LoginDetails struct {
	Email    string
	Password string
}

type ErrorMessage struct {
	ErrorMessage string `json:"ErrorMessage"`
}
