package status

type Store interface {
	SaveSubscription(profile Profile) error
	DeleteSubscription(profile Profile) error
	UpdateSubscription(status *ProfileStatus) error
	GetSubscription(url string) (*ProfileStatus, error)

	SaveSubscriber(profile Profile) error
	RemoveSubscriber(profile Profile) error

	SaveStatusItem(item StatusItem, ownerid string) error
	GetStatus() ([]*ProfileStatus, error)

	GetSubscriber() ([]*Profile, error)
}