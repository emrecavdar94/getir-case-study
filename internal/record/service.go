package record

import (
	"go.uber.org/zap"
	"time"
)

type (
	Repository interface {
		GetRecordsByDateAndCount(startDate, endDate time.Time, minCount, maxCount int) ([]*Record, error)
	}

	Service struct {
		repository Repository
		logger     *zap.Logger
	}
)

func NewService(repository Repository, logger *zap.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetRecordsByDateAndCount(recordRequest RecordRequest) ([]*Record, error) {
	s.logger.Sugar().Debugf("Getting records by %+v", recordRequest)

	t, err := time.Parse("2006-01-02", recordRequest.StartDate)

	if err != nil {
		return nil, err
	}
	t2, err := time.Parse("2006-01-02", recordRequest.EndDate)

	if err != nil {
		return nil, err
	}

	records, err := s.repository.GetRecordsByDateAndCount(t, t2, recordRequest.MinCount, recordRequest.MaxCount)
	if err != nil {
		s.logger.Sugar().Errorf("error while getting records by startDate: %s, endDate:%s, minCount: %d, maxCount: %d", recordRequest.StartDate, recordRequest.EndDate, recordRequest.MinCount, recordRequest.MaxCount)
		return records, err
	}
	s.logger.Sugar().Debugf("Records successfully got")
	return records, err
}
