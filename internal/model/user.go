package models

import "github.com/volatiletech/null"

// User struct
type User struct {
	ID       uint64      `json:"id"`
	FullName string      `json:"full_name"`
	UserName string      `json:"username"`
	Password string      `json:"password"`
	Gender   null.String `json:"gender"`
	Hobby    null.String `json:"hobby"`
	Address  null.String `json:"address"`
}
