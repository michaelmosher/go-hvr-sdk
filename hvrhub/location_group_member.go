package hvrhub

type LocationGroupMember struct {
	ChannelName       string `db:"chn_name"`
	LocationGroupName string `db:"grp_name"`
	LocationName      string `db:"loc_name"`
}
