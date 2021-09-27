package topic

type Topic struct {
	TopicID     int    `json:"topicID"`
	CommunityID int    `json:"communityID"`
	CreatedDate string `json:"createdDate"`
	CreatedBy   string `json:"createdBy"`
	TopicName   string `json:"topicName"`
}
