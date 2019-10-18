package status

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"
)

type loggingService struct {
	logger  log.Logger
	service Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger: logger, service: s}
}


func (s *loggingService) AddSubscriber(profile Profile) error {
	start := time.Now()

	err := s.service.AddSubscriber(profile)

	logger := log.With(s.logger,
		"method", "addSubscriber",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to add subscriber", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return err
}

func (s *loggingService) RemoveSubscriber(profile Profile) error {
	start := time.Now()

	err := s.service.RemoveSubscriber(profile)

	logger := log.With(s.logger,
		"method", "removeSubscriber",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to remove subscriber", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return err
}

func (s *loggingService) AddStatus(status StatusItem) error {
	start := time.Now()

	err := s.service.AddStatus(status)

	logger := log.With(s.logger,
		"method", "addStatus",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to add status item", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return err
}

func (s *loggingService) GetStatus() ([]*ProfileStatus, error) {
	start := time.Now()

	stat, err := s.service.GetStatus()

	logger := log.With(s.logger,
		"method", "getStatus",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to get status", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return stat, err
}

func (s *loggingService) SubscribeTo(profile Profile) error {
	start := time.Now()

	err := s.service.SubscribeTo(profile)

	logger := log.With(s.logger,
		"method", "addSubscription",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to add subscription", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return err
}

func (s *loggingService) UnsubscribeFrom(profile Profile) error {
	start := time.Now()

	err := s.service.UnsubscribeFrom(profile)

	logger := log.With(s.logger,
		"method", "removeSubscription",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to remove subscription", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return err
}

func (s *loggingService) UpdateSubscriptionFrom(status ProfileStatus) error {
	start := time.Now()

	err := s.service.UpdateSubscriptionFrom(status)

	logger := log.With(s.logger,
		"method", "notify",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log("msg", "failed to update status", "err", err)
	} else {
		level.Debug(logger).Log()
	}

	return err
}


