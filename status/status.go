package status

type Profile struct {
	URL string `json:"url"`
}

type ProfileStatus struct {
	URL    string       `json:"url"`
	Alias  string       `json:"alias"`
	Status []StatusItem `json:"status"`
}

type StatusItem struct {
	ID      string `json:"id"`
	Media   string `json:"media"`
	Caption string `json:"caption"`
}

type Service interface {
	// subscribers
	AddSubscriber(profile Profile) error
	RemoveSubscriber(profile Profile) error

	// our status
	AddStatus(status StatusItem) error

	// our + subscribed status
	// called by the client
	GetStatus() ([]*ProfileStatus, error)

	// our subscriptions
	// called by the client
	SubscribeTo(profile Profile) error
	UnsubscribeFrom(profile Profile) error

	// receive updates
	UpdateSubscriptionFrom(status ProfileStatus) error
}

type service struct {
	OwnerUrl string
	Store    Store
}

func NewService(publicAddr string, alias string) *service {
	return &service{
		OwnerUrl: publicAddr,
		Store: NewFakeStore(publicAddr, alias),
	}
}

func (s *service) AddSubscriber(profile Profile) error {
	return s.Store.SaveSubscriber(profile)
}

func (s *service) RemoveSubscriber(profile Profile) error {
	return s.Store.RemoveSubscriber(profile)
}

func (s *service) AddStatus(status StatusItem) error {
	err := s.Store.SaveStatusItem(status, s.OwnerUrl)
	if err != nil {
		return err
	}
	ownerstatus, err := s.Store.GetSubscription(s.OwnerUrl)
	if err != nil {
		return err
	}
	subscriber, err :=s.Store.GetSubscriber()
	if err != nil {
		return err
	}
	return Notify(*ownerstatus, subscriber)
}

func (s *service) GetStatus() ([]*ProfileStatus, error) {
	return s.Store.GetStatus()
}

func (s *service) SubscribeTo(profile Profile) error {
	err := AddSubscription(Profile{URL: s.OwnerUrl}, profile.URL)
	if err != nil {
		return err
	}
	return s.Store.SaveSubscription(profile)

}

func (s *service) UnsubscribeFrom(profile Profile) error {
	err := RemoveSubscription(Profile{URL: s.OwnerUrl}, profile.URL)
	if err != nil {
		return err
	}
	return s.Store.DeleteSubscription(profile)
}

func (s *service) UpdateSubscriptionFrom(status ProfileStatus) error {
	return s.Store.UpdateSubscription(&status)
}
