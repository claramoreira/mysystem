package topic

import (
	"database/sql"
	"system/database"
)

func getTopicList() ([]Topic, error) {
	results, err := database.DbConn.Query(`
	SELECT topic_id, community_id, created_date, created_by, topic_name
	FROM tbTopic`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	topics := make([]Topic, 0)
	for results.Next() {
		var topic Topic
		results.Scan(&topic.TopicID,
			&topic.CommunityID,
			&topic.CreatedDate,
			&topic.CreatedBy,
			&topic.TopicName)
		topics = append(topics, topic)
	}
	return topics, nil
}

func getTopic(topicID int) (*Topic, error) {
	row := database.DbConn.QueryRow(`
	SELECT topic_id, community_id, created_date, created_by, topic_name
	FROM tbTopic
	WHERE topic_id = ?`, topicID)
	topic := &Topic{}
	err := row.Scan(&topic.TopicID,
		&topic.CommunityID,
		&topic.CreatedDate,
		&topic.CreatedBy,
		&topic.TopicName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return topic, nil
}

func getTopicByCommunity(communityID int) ([]Topic, error) {
	results, err := database.DbConn.Query(`
	SELECT topic_id, community_id, created_date, created_by, topic_name
	FROM tbTopic
	WHERE community_id = ?`, communityID)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	topics := make([]Topic, 0)
	for results.Next() {
		var topic Topic
		results.Scan(&topic.TopicID,
			&topic.CommunityID,
			&topic.CreatedDate,
			&topic.CreatedBy,
			&topic.TopicName)
		topics = append(topics, topic)
	}
	return topics, nil
}

func removeTopic(topicID int) error {
	_, err := database.DbConn.Query(`DELETE FROM tbTopic where topic_id = ?`,
		topicID)
	if err != nil {
		return err
	}
	return nil
}

func insertTopic(topic Topic) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO tbTopic
	(community_id, created_date, created_by, topic_name)
	VALUES
	(?, sysdate(), ?, ?)`,
		topic.CommunityID, topic.CreatedBy, topic.TopicName)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}
