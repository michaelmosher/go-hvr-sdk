package hvrhub

type Channel struct {
	Name        string `db:"chn_name"`
	Description string `db:"chn_description"`
}
