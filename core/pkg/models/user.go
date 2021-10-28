package models

import "encoding/json"

type User struct {
	Email string `json:"email"`
}

func UserFromJson(userJson string) (*User, error) {
	user := User{}
	err := json.Unmarshal([]byte(userJson), &user)

	return &user, err
}
