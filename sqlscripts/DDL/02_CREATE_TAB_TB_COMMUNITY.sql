CREATE TABLE go.tbCommunity (
	community_id INT NOT NULL AUTO_INCREMENT,
	created_date DATETIME NOT NULL,
	updated_date DATETIME,
	created_by INT NOT NULL,
	community_name VARCHAR(80) NOT NULL UNIQUE,
	community_description TEXT NOT NULL,
	community_avatar VARCHAR(100),
	PRIMARY KEY (community_id),
	FOREIGN KEY (created_by) REFERENCES go.tbUser(user_id)
);