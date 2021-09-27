package user

import "database/sql"

type User struct {
	UserID          int            `json:"userID"`
	UserName        string         `json:"userName"`
	UserUsername    string         `json:"userUsername"`
	UserEmail       string         `json:"userEmail"`
	UserAvatar      sql.NullString `json:"userAvatar"`
	UserDescription sql.NullString `json:"userDescription"`
	CreatedDate     string         `json:"createdDate"`
}
