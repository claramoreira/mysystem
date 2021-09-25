package user

type User struct {
	UserID          int    `json:"userID"`
	UserName        string `json:"userName"`
	UserUsername    string `json:"userUsername"`
	UserEmail       string `json:"userEmail"`
	UserAvatar      string `json:"userAvatar"`
	UserDescription string `json:"userDescription"`
	CreatedDate     string `json:"createdDate"`
}
