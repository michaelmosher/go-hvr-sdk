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
