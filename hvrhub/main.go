package hvrhub

type client interface {
	GetLocation(string) (Location, error)
	NewLocation(Location) error
	UpdateLocation(Location) error
	DeleteLocation(string) error

	GetChannel(string) (Channel, error)
	NewChannel(Channel) error
	UpdateChannel(Channel) error
	DeleteChannel(string) error

	GetLocationGroup(string, string) (LocationGroup, error)
	NewLocationGroup(LocationGroup) error
	UpdateLocationGroup(LocationGroup) error
	DeleteLocationGroup(string, string) error
}

type Service struct {
	Client client
}

func (s Service) GetLocation(locationName string) (Location, error) {
	return s.Client.GetLocation(locationName)
}

func (s Service) NewLocation(l Location) error {
	return s.Client.NewLocation(l)
}

func (s Service) UpdateLocation(l Location) error {
	return s.Client.UpdateLocation(l)
}

func (s Service) DeleteLocation(locationName string) error {
	return s.Client.DeleteLocation(locationName)
}

func (s Service) GetChannel(channelName string) (Channel, error) {
	return s.Client.GetChannel(channelName)
}

func (s Service) NewChannel(c Channel) error {
	return s.Client.NewChannel(c)
}

func (s Service) UpdateChannel(c Channel) error {
	return s.Client.UpdateChannel(c)
}

func (s Service) DeleteChannel(channelName string) error {
	return s.Client.DeleteChannel(channelName)
}

func (s Service) GetLocationGroup(channelName, groupName string) (LocationGroup, error) {
	return s.Client.GetLocationGroup(channelName, groupName)
}

func (s Service) NewLocationGroup(g LocationGroup) error {
	return s.Client.NewLocationGroup(g)
}

func (s Service) UpdateLocationGroup(g LocationGroup) error {
	return s.Client.UpdateLocationGroup(g)
}

func (s Service) DeleteLocationGroup(channelName, groupName string) error {
	return s.Client.DeleteLocationGroup(channelName, groupName)
}
