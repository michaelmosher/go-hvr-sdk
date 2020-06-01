package postgresql

import (
	"fmt"

	"github.com/michaelmosher/go-hvr-sdk/hvrhub"
)

func (s service) GetLocation(locationName string) (hvrhub.Location, error) {
	var location hvrhub.Location
	err := s.db.Get(&location, "SELECT * FROM hvr_location WHERE loc_name = $1", locationName)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return location, fmt.Errorf("location not found")
		}
		return location, fmt.Errorf("error getting location %s: %s", locationName, err)
	}

	return location, nil
}

func (s service) NewLocation(l hvrhub.Location) error {
	insertStmt := `INSERT INTO hvr_location(
		loc_name,
		loc_class,
		loc_directory,
		loc_remote_node,
		loc_remote_login,
		loc_remote_pwd,
		loc_remote_port,
		loc_db_name,
		loc_db_user,
		loc_description
	) VALUES (
		:loc_name,
		:loc_class,
		:loc_directory,
		:loc_remote_node,
		:loc_remote_login,
		:loc_remote_pwd,
		:loc_remote_port,
		:loc_db_name,
		:loc_db_user,
		:loc_description
	)`
	_, err := s.db.NamedExec(insertStmt, l)
	return err
}

func (s service) UpdateLocation(l hvrhub.Location) error {
	updateStmt := `UPDATE hvr_location SET
		loc_name = :loc_name,
		loc_class = :loc_class,
		loc_directory = :loc_directory,
		loc_remote_node = :loc_remote_node,
		loc_remote_login = :loc_remote_login,
		loc_remote_pwd = :loc_remote_pwd,
		loc_remote_port = :loc_remote_port,
		loc_db_name = :loc_db_name,
		loc_db_user = :loc_db_user,
		loc_description = :loc_description
	WHERE loc_name = :loc_name`
	_, err := s.db.NamedExec(updateStmt, l)
	return err
}

func (s service) DeleteLocation(l hvrhub.Location) error {
	deleteStmt := "DELETE from hvr_location WHERE loc_name = :loc_name"
	_, err := s.db.NamedExec(deleteStmt, l)
	return err
}
