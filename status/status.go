package status

import (
	"fmt"
)

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
	Owner         ProfileStatus
	State         map[string]ProfileStatus
	Subscribers   map[string]Profile
	Subscriptions map[string]Profile
}

func NewService(publicAddr string, alias string) *service {
	return &service{
		Owner: ProfileStatus{
			URL:    publicAddr,
			Alias:  alias,
			Status: []StatusItem{},
		},
		State:         map[string]ProfileStatus{},
		Subscribers:   map[string]Profile{},
		Subscriptions: map[string]Profile{},
	}
}

func (s *service) AddSubscriber(profile Profile) error {
	s.Subscribers[profile.URL] = profile
	return nil
}

func (s *service) RemoveSubscriber(profile Profile) error {
	delete(s.Subscribers, profile.URL)
	return nil
}

func (s *service) AddStatus(status StatusItem) error {
	s.Owner.Status = append(s.Owner.Status, status)
	err := Notify(s.Owner, s.Subscribers)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetStatus() ([]*ProfileStatus, error) {
	state := make([]*ProfileStatus, 0, len(s.State))
	for _, profileStatus := range s.State {
		state = append(state, &profileStatus)
	}
	return state, nil
}

func (s *service) SubscribeTo(profile Profile) error {
	err := AddSubscription(Profile{URL: s.Owner.URL}, profile.URL)
	if err != nil {
		return err
	}
	s.Subscriptions[profile.URL] = profile
	return nil
}

func (s *service) UnsubscribeFrom(profile Profile) error {
	err := RemoveSubscription(Profile{URL: s.Owner.URL}, profile.URL)
	if err != nil {
		return err
	}
	delete(s.Subscriptions, profile.URL)
	return nil
}

func (s *service) UpdateSubscriptionFrom(status ProfileStatus) error {
	if _, subscribed := s.Subscriptions[status.URL]; !subscribed {
		return fmt.Errorf("no subscription for %s", status.URL)
	}
	s.State[status.URL] = status
	return nil
}
