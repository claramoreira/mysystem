CREATE TABLE go.tbPost (
	post_id INT NOT NULL AUTO_INCREMENT,
	topic_id INT NOT NULL,
	created_date DATETIME NOT NULL,
	created_by INT NOT NULL,
	post_content TEXT CHARACTER SET Latin1  COLLATE latin1_general_ci NOT NULL,
	PRIMARY KEY (post_id),
	FOREIGN KEY (created_by) REFERENCES go.tbUser(user_id),
    FOREIGN KEY (topic_id) REFERENCES go.tbTopic(topic_id)
);