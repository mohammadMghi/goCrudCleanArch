package domain

type User struct{
	ID int64 `json:"id"`
	Phone string `json:"phone"`
	Name string `json:"name"`
	Password string `json:"password"`
}

