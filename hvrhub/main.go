package hvrhub

type client interface {
	GetLocation(string) (Location, error)
	NewLocation(Location) error
	UpdateLocation(Location) error
	DeleteLocation(Location) error
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

func (s Service) DeleteLocation(l Location) error {
	return s.Client.DeleteLocation(l)
}
