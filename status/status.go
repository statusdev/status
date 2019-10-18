package status

type Status struct {
	ID      string `json:"id"`
	Media   string `json:"media"`
	Caption string `json:"caption"`
}

type User struct {
	URL    string   `json:"url"`
	Alias  string   `json:"alias"`
	Status []string `json:"status"`
}

type Service interface {
	// subscribers
	AddSubscriber(user User) error
	RemoveSubscriber(user User) error

	// our status
	AddStatus(status Status) error
	GetStatus() ([]Status, error)

	// our subscriptions
	SubscribeTo(user User) error
	UnsubscribeTo(user User) error

	// receive updates
	UpdateSubscriptionFrom(user User) error
	GetSubscriptions() ([]User, error)
}

type service struct {
	Owner         User
	Subscribers   []User
	Subscriptions []User
}

func NewService() *service {
	return new(service)
}

func (s *service) Subscribe(subscriber User) error {
	// implement me
	return nil
}

func (s *service) Unsubscribe(subscriber User) error {
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
