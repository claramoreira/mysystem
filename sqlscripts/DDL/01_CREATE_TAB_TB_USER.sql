CREATE TABLE go.tbUser (
	user_id INT NOT NULL AUTO_INCREMENT,
	user_name VARCHAR(50) NOT NULL UNIQUE,
	user_username VARCHAR(20) NOT NULL UNIQUE,
	user_email VARCHAR(50) NOT NULL UNIQUE,
	user_password VARCHAR(20) NOT NULL,
	user_avatar VARCHAR(100),
	user_description TEXT,
	created_date DATETIME NOT NULL,
	updated_date DATETIME,
	PRIMARY KEY (user_id)
);