package user

import "system/database"

func getUsersList() ([]User, error) {
	results, err := database.DbConn.Query(`
	SELECT user_id, user_name, user_username, user_email, user_avatar, user_description, created_date
	FROM tbUser`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	users := make([]User, 0)
	for results.Next() {
		var user User
		results.Scan(&user.UserID,
			&user.UserName,
			&user.UserUsername,
			&user.UserEmail,
			&user.UserAvatar,
			&user.UserDescription,
			&user.CreatedDate)
		users = append(users, user)
	}
	return users, nil
}
