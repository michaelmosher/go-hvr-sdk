package postgresql

import (
	"fmt"

	"github.com/michaelmosher/go-hvr-sdk/hvrhub"
)

func (s service) GetLocationGroup(channelName, groupName string) (hvrhub.LocationGroup, error) {
	var group hvrhub.LocationGroup

	selectStmt := `SELECT * FROM hvr_loc_group WHERE chn_name = $1 and grp_name = $2`
	err := s.db.Get(&group, selectStmt, channelName, groupName)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return group, fmt.Errorf("location group not found")
		}
		return group, fmt.Errorf("error getting location group (%s %s): %s", channelName, groupName, err)
	}

	return group, nil
}

func (s service) FindLocationGroups(channelName string) ([]hvrhub.LocationGroup, error) {
	groups := []hvrhub.LocationGroup{}

	selectStmt := `SELECT * from hvr_loc_group WHERE chn_name = $1`
	if err := s.db.Select(&groups, selectStmt, channelName); err != nil {
		return groups, fmt.Errorf("error finding location groups (%s): %s", channelName, err)
	}

	return groups, nil
}

func (s service) NewLocationGroup(g hvrhub.LocationGroup) error {
	insertStmt := `INSERT INTO hvr_loc_group(
		chn_name,
		grp_name,
		grp_description
	) VALUES (
		:chn_name,
		:grp_name,
		:grp_description
	)`
	_, err := s.db.NamedExec(insertStmt, g)
	return err
}

func (s service) UpdateLocationGroup(g hvrhub.LocationGroup) error {
	updateStmt := `UPDATE hvr_loc_group SET
		grp_description = :grp_description
	WHERE chn_name = :chn_name
		and grp_name = :grp_name`
	_, err := s.db.NamedExec(updateStmt, g)
	return err
}

func (s service) DeleteLocationGroup(channelName, groupName string) error {
	deleteStmt := "DELETE FROM hvr_loc_group WHERE chn_name = $1 and grp_name = $2"
	_, err := s.db.Exec(deleteStmt, channelName, groupName)
	return err
}
