package post

type Post struct {
	PostID      int    `json:"postID"`
	TopicID     int    `json:"topicID"`
	CreatedDate string `json:"createdDate"`
	CreatedBy   int    `json:"createdBy"`
	PostContent string `json:"postContent"`
}
