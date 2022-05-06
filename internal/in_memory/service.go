package inmemory

import "go.uber.org/zap"

type Repository interface {
	Get(key string) (*InMemory, error)
	Set(key string, value string) (*InMemory, error)
}

type Service struct {
	Repository Repository
	logger     *zap.Logger
}

func NewService(repository Repository, logger *zap.Logger) *Service {
	return &Service{
		Repository: repository,
		logger:     logger,
	}
}

func (s *Service) Get(key string) (*InMemory, error) {
	s.logger.Sugar().Debugf("Get key: %s", key)
	data, err := s.Repository.Get(key)
	if err != nil {
		s.logger.Sugar().Errorf("Get key: %s, error: %s", key, err.Error())
		return data, err
	}
	s.logger.Sugar().Debugf("Get key: %s, value: %s", key, data.Value)
	return data, err
}

func (s *Service) Set(key string, value string) (*InMemory, error) {
	s.logger.Sugar().Debugf("Set key: %s, value: %s", key, value)
	data, err := s.Repository.Set(key, value)
	if err != nil {
		s.logger.Sugar().Errorf("Set key: %s, value: %s, error: %s", key, value, err.Error())
		return data, err
	}
	s.logger.Sugar().Debugf("Set key: %s, value: %s. Completed successfully", key, value)
	return data, err
}
