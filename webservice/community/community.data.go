package community

import (
	"database/sql"
	"system/database"
)

func getCommunitiesList() ([]Community, error) {
	results, err := database.DbConn.Query(`
	SELECT community_id, created_date, created_by, updated_date, community_name, community_description, community_avatar
	FROM tbCommunity`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	communities := make([]Community, 0)
	for results.Next() {
		var community Community
		results.Scan(&community.CommunityID,
			&community.CreatedDate,
			&community.CreatedBy,
			&community.UpdatedDate,
			&community.CommunityName,
			&community.CommunityDescription,
			&community.CommunityAvatar)
		communities = append(communities, community)
	}
	return communities, nil
}

func getCommunity(communityID int) (*Community, error) {
	row := database.DbConn.QueryRow(`
	SELECT community_id, created_date, created_by, updated_date, community_name, community_description, community_avatar
	FROM tbCommunity
	WHERE community_id = ?`, communityID)
	community := &Community{}
	err := row.Scan(&community.CommunityID,
		&community.CreatedDate,
		&community.CreatedBy,
		&community.UpdatedDate,
		&community.CommunityName,
		&community.CommunityDescription,
		&community.CommunityAvatar)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return community, nil
}

func removeCommunity(communityID int) error {
	_, err := database.DbConn.Query(`DELETE FROM tbCommunity where community_id = ?`, communityID)
	if err != nil {
		return err
	}
	return nil
}

func insertCommunity(community Community) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO tbCommunity
	(created_date, created_by, community_name, community_description, community_avatar)
	VALUES
	(sysdate(), ?, ?, ?, ?)`,
		community.CreatedBy, community.CommunityName, community.CommunityDescription, community.CommunityAvatar)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}

func updateCommunity(community Community) error {
	_, err := database.DbConn.Exec(`UPDATE tbCommunity SET
	community_name = ?,
	community_description = ?,
	community_avatar = ?,
	updated_date = sysdate()
	WHERE community_id=?`,
		community.CommunityName,
		community.CommunityDescription,
		community.CommunityAvatar,
		community.CommunityID)
	if err != nil {
		return err
	}
	return nil
}
