package post

import (
	"database/sql"
	"system/database"
)

func getPostList() ([]Post, error) {
	results, err := database.DbConn.Query(`
	SELECT post_id, topic_id, created_date, created_by, post_content
	FROM tbPost`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	posts := make([]Post, 0)
	for results.Next() {
		var post Post
		results.Scan(&post.PostID,
			&post.TopicID,
			&post.CreatedDate,
			&post.CreatedBy,
			&post.PostContent)
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(postID int) (*Post, error) {
	row := database.DbConn.QueryRow(`
	SELECT post_id, topic_id, created_date, created_by, post_content
	FROM tbPost
	WHERE post_id = ?`, postID)
	post := &Post{}
	err := row.Scan(&post.PostID,
		&post.TopicID,
		&post.CreatedDate,
		&post.CreatedBy,
		&post.PostContent)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return post, nil
}

func removePost(postID int) error {
	_, err := database.DbConn.Query(`DELETE FROM tbPost where post_id = ?`,
		postID)
	if err != nil {
		return err
	}
	return nil
}

func insertPost(post Post) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO tbPost
	(topic_id, created_date, created_by, post_content)
	VALUES
	(?, sysdate(), ?, ?)`,
		post.TopicID, post.CreatedBy, post.PostContent)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}
