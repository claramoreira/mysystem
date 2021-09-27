package user

import "system/utils"

type User struct {
	UserID          int                  `json:"userID"`
	UserName        string               `json:"userName"`
	UserUsername    string               `json:"userUsername"`
	UserEmail       string               `json:"userEmail"`
	UserPassword    string               `json:"userPassword"`
	UserAvatar      utils.JSONNullString `json:"userAvatar"`
	UserDescription utils.JSONNullString `json:"userDescription"`
	CreatedDate     string               `json:"createdDate"`
}
