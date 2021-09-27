package user

import (
	"database/sql"
	"system/database"
)

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

func getUser(userID int) (*User, error) {
	row := database.DbConn.QueryRow(`
	SELECT user_id, user_name, user_username, user_email, user_avatar, user_description, created_date
	FROM tbUser
	WHERE user_id = ?`, userID)
	user := &User{}
	err := row.Scan(&user.UserID,
		&user.UserName,
		&user.UserUsername,
		&user.UserEmail,
		&user.UserAvatar,
		&user.UserDescription,
		&user.CreatedDate)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func removeUser(userID int) error {
	_, err := database.DbConn.Query(`DELETE FROM tbUser where user_id = ?`, userID)
	if err != nil {
		return err
	}
	return nil
}

func insertUser(user User) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO tbUser
	(user_name, user_username, user_email, user_password, created_date)
	VALUES
	(?, ?, ?, ?, sysdate())`,
		user.UserName, user.UserUsername, user.UserEmail, user.UserPassword)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}

func updateUser(user User) error {
	_, err := database.DbConn.Exec(`UPDATE tbUser SET
	user_name = ?,
	user_username = ?,
	user_email = ?,
	user_password = ?,
	user_avatar = ?,
	user_description = ?,
	updated_date = sysdate()
	WHERE user_id=?`,
		user.UserName,
		user.UserUsername,
		user.UserEmail,
		user.UserPassword,
		user.UserAvatar,
		user.UserDescription,
		user.UserID)
	if err != nil {
		return err
	}
	return nil
}
