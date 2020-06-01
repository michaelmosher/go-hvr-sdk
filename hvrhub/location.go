package hvrhub

type Location struct {
	Name           string `db:"loc_name"`
	Class          string `db:"loc_class"`
	Directory      string `db:"loc_directory"`
	RemoteNode     string `db:"loc_remote_node"`
	RemoteLogin    string `db:"loc_remote_login"`
	RemotePassword string `db:"loc_remote_pwd"`
	RemotePort     int    `db:"loc_remote_port"`
	DBName         string `db:"loc_db_name"`
	DBUser         string `db:"loc_db_user"`
	Description    string `db:"loc_description"`
}
