package hvrhub

type LocationGroup struct {
	ChannelName string `db:"chn_name"`
	Name        string `db:"grp_name"`
	Description string `db:"grp_description"`
}
