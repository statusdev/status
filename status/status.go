package status

type Status struct {
	ID      string `json:"id"`
	Media   string `json:"media"`
	Caption string `json:"caption"`
}

type Subscriber struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

type Service interface {
	Subscribe(subscriber Subscriber) error
	Unsubscribe(subscriber Subscriber) error
	AddStatus(status Status) error
	GetStatus() ([]Status, error)
	AddNotification(status Status) error
	GetNotifications() ([]Status, error)
}

type service struct {
	Subscribers   []Subscriber
	Status        []Status
	Notifications []Status
}

func NewService() *service {
	return new(service)
}

func (s *service) Subscribe(subscriber Subscriber) error {
	// implement me
	return nil
}

func (s *service) Unsubscribe(subscriber Subscriber) error {
	// implement me
	return nil
}

func (s *service) AddStatus(status Status) error {
	// implement me
	return nil
}

func (s *service) GetStatus() ([]Status, error) {
	// implement me
	return []Status{}, nil
}

func (s *service) AddNotifications(status Status) error {
	// implement me
	return nil
}

func (s *service) GetNotifications() ([]Status, error) {
	// implement me
	return []Status{}, nil
}
