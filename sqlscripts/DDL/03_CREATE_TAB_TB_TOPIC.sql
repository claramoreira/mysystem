	CREATE TABLE go.tbTopic (
		topic_id INT NOT NULL AUTO_INCREMENT,
		community_id INT NOT NULL,
		created_date DATETIME NOT NULL,
		created_by INT NOT NULL,
		topic_name VARCHAR(80) NOT NULL,
		PRIMARY KEY (topic_id),
		FOREIGN KEY (created_by) REFERENCES go.tbUser(user_id),
		FOREIGN KEY (community_id) REFERENCES go.tbCommunity(community_id)
	);