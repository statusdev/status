package status

import "fmt"

type Fakestore struct {
	Subscribers   map[string]*Profile
	Subscriptions map[string]*ProfileStatus
}

func NewFakeStore(ownerid string, owneralias string) Store {
	return Fakestore{
		Subscribers: map[string]*Profile{},
		Subscriptions: map[string]*ProfileStatus{
			ownerid: {
				URL:    ownerid,
				Alias:  owneralias,
				Status: []StatusItem{},
			},
		},
	}
}

func (f Fakestore) SaveSubscription(profile Profile) error {
	f.Subscriptions[profile.URL] = &ProfileStatus{
		URL:    profile.URL,
		Alias:  "",
		Status: []StatusItem{},
	}
	return nil
}

func (f Fakestore) DeleteSubscription(profile Profile) error {
	delete(f.Subscriptions, profile.URL)
	return nil
}

func (f Fakestore) UpdateSubscription(status *ProfileStatus) error {
	if _, exists := f.Subscriptions[status.URL]; !exists {
		return fmt.Errorf("no subscription with url %s", status.URL)
	}
	f.Subscriptions[status.URL] = status
	return nil
}

func (f Fakestore) SaveSubscriber(profile Profile) error {
	f.Subscribers[profile.URL] = &profile
	return nil
}

func (f Fakestore) RemoveSubscriber(profile Profile) error {
	delete(f.Subscribers, profile.URL)
	return nil
}

func (f Fakestore) SaveStatusItem(item StatusItem, ownerid string) error {
	status := f.Subscriptions[ownerid].Status
	status = append(status, item)
	f.Subscriptions[ownerid].Status = status
	return nil
}

func (f Fakestore) GetStatus() ([]*ProfileStatus, error) {
	status := make([]*ProfileStatus, 0, len(f.Subscriptions))
	for _, profile := range f.Subscriptions {
		status = append(status, profile)
	}
	return status, nil
}

func (f Fakestore) GetSubscription(url string) (*ProfileStatus, error) {
	status, exists := f.Subscriptions[url]
	if !exists {
		return nil, fmt.Errorf("no subscription for %s", url)
	}
	return status, nil
}

func (f Fakestore) GetSubscriber() ([]*Profile, error) {
	subscriber := make([]*Profile, 0, len(f.Subscribers))
	for _, sub := range f.Subscribers {
		subscriber = append(subscriber, sub)
	}
	return subscriber, nil
}
