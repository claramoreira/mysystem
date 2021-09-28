package community

import "system/utils"

type Community struct {
	CommunityID          int                  `json:"communityID"`
	CreatedDate          string               `json:"createdDate"`
	CreatedBy            int                  `json:"createdBy"`
	UpdatedDate          utils.JSONNullString `json:"updatedDate"`
	CommunityName        string               `json:"communityName"`
	CommunityDescription string               `json:"communityDescription"`
	CommunityAvatar      utils.JSONNullString `json:"communityAvatar"`
}
