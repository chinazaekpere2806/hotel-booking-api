package models

type Guest struct {
	ID          string `json:"id"`
	Firstname   string `json:"firstName"`
	Lastname    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
}
