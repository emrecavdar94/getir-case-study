package record_test

import (
	"fmt"
	"getir-assignment/internal/record"
	"getir-assignment/internal/record/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestService_GetRecordsByDateAndCount(t *testing.T) {
	mockRepository := createMockRepository(t)
	service := record.NewService(mockRepository, zap.NewNop())
	t.Run("Given  recordRequest when GetRecordsByDateAndCount called it should return recordData", func(t *testing.T) {
		recordRequest := record.RecordRequest{
			StartDate: "2021-05-05",
			EndDate:   "2021-05-06",
			MinCount:  10,
			MaxCount:  20,
		}
		sd, _ := time.Parse("2006-01-02", recordRequest.StartDate)
		ed, _ := time.Parse("2006-01-02", recordRequest.EndDate)
		expectedRecords := []*record.Record{
			{
				Key:       "1111",
				CreatedAt: time.Now().Add(-10 * time.Hour),
				Counts:    100,
			},
		}
		mockRepository.
			EXPECT().
			GetRecordsByDateAndCount(sd,
				ed,
				recordRequest.MinCount,
				recordRequest.MaxCount).
			Return(expectedRecords, nil).Times(1)

		data, err := service.GetRecordsByDateAndCount(recordRequest)
		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
	t.Run("Given recordRequest with incorrect data when GetRecordsByDateAndCount called it should return error", func(t *testing.T) {
		recordRequest := record.RecordRequest{
			StartDate: "2021-05-05",
			EndDate:   "2021-05-06",
			MinCount:  10,
			MaxCount:  20,
		}
		sd, _ := time.Parse("2006-01-02", recordRequest.StartDate)
		ed, _ := time.Parse("2006-01-02", recordRequest.EndDate)
		mockRepository.
			EXPECT().
			GetRecordsByDateAndCount(sd,
				ed,
				recordRequest.MinCount,
				recordRequest.MaxCount).
			Return(nil, fmt.Errorf("test")).Times(1)

		data, err := service.GetRecordsByDateAndCount(recordRequest)
		assert.Nil(t, data)
		assert.NotNil(t, err)
	})
}

func createMockRepository(t *testing.T) *mocks.MockRepository {
	c := gomock.NewController(t)
	repository := mocks.NewMockRepository(c)
	return repository
}
