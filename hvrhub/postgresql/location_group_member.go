package postgresql

import (
	"fmt"

	"github.com/michaelmosher/go-hvr-sdk/hvrhub"
)

func (s service) GetLocationGroupMember(channelName, groupName, locationName string) (hvrhub.LocationGroupMember, error) {
	var group hvrhub.LocationGroupMember

	selectStmt := `SELECT * FROM hvr_loc_group_member
	WHERE chn_name = $1
		AND grp_name = $2
		AND loc_name = $3`
	err := s.db.Get(&group, selectStmt, channelName, groupName, locationName)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return group, fmt.Errorf("location group member not found")
		}
		return group, fmt.Errorf("error getting location group member (%s %s %s): %s",
			channelName, groupName, locationName, err)
	}

	return group, nil
}

func (s service) FindLocationGroupMembers(channelName, groupName string) ([]hvrhub.LocationGroupMember, error) {
	members := []hvrhub.LocationGroupMember{}

	selectStmt := `SELECT * from hvr_loc_group_member WHERE chn_name = $1 AND grp_name = $2`
	if err := s.db.Select(&members, selectStmt, channelName, groupName); err != nil {
		return members, fmt.Errorf("error finding location group members (%s %s): %s", channelName, groupName, err)
	}

	return members, nil
}

func (s service) NewLocationGroupMember(gm hvrhub.LocationGroupMember) error {
	insertStmt := `INSERT INTO hvr_loc_group_member(
		chn_name,
		grp_name,
		loc_name
	) VALUES (
		:chn_name,
		:grp_name,
		:loc_name
	)`
	_, err := s.db.NamedExec(insertStmt, gm)
	return err
}

func (s service) DeleteLocationGroupMember(channelName, groupName, locationName string) error {
	deleteStmt := `DELETE FROM hvr_loc_group_member
	WHERE chn_name = $1
		AND grp_name = $2
		AND loc_name = $3`
	_, err := s.db.Exec(deleteStmt, channelName, groupName, locationName)
	return err
}
