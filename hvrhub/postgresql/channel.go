package postgresql

import (
	"fmt"

	"github.com/michaelmosher/go-hvr-sdk/hvrhub"
)

func (s service) GetChannel(channelName string) (hvrhub.Channel, error) {
	var channel hvrhub.Channel
	err := s.db.Get(&channel, "SELECT * FROM hvr_channel WHERE chn_name = $1", channelName)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return channel, fmt.Errorf("channel not found")
		}
		return channel, fmt.Errorf("error getting channel %s: %s", channelName, err)
	}

	return channel, nil
}

func (s service) NewChannel(c hvrhub.Channel) error {
	insertStmt := `INSERT INTO hvr_channel(
		chn_name,
		chn_description
	) VALUES (
		:chn_name,
		:chn_description
	)`
	_, err := s.db.NamedExec(insertStmt, c)
	return err
}

func (s service) UpdateChannel(c hvrhub.Channel) error {
	updateStmt := `UPDATE hvr_channel SET
		chn_description = :chn_description
	WHERE chn_name = :chn_name`
	_, err := s.db.NamedExec(updateStmt, c)
	return err
}

func (s service) DeleteChannel(channelName string) error {
	deleteStmt := "DELETE from hvr_channel WHERE chn_name = $1"
	_, err := s.db.Exec(deleteStmt, channelName)
	return err
}
