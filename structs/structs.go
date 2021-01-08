package structs

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FullName string `json:"fullName"`
	Username string `json:"userName"`
	Phone    int    `json:"userName"`
	Address  string `json:"address"`
}
