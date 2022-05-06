package record_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"getir-assignment/internal/record"
	"getir-assignment/internal/record/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHandler_GetRecordsByDateAndCount(t *testing.T) {
	mockService := createMockService(t)

	handler := record.NewHandler(mockService, zap.NewNop())
	t.Run("successful", func(t *testing.T) {
		recordRequest := record.RecordRequest{
			StartDate: "2021-05-05",
			EndDate:   "2021-05-06",
			MinCount:  10,
			MaxCount:  20,
		}
		expectedRecords := []*record.Record{
			{
				Key:       "1111",
				CreatedAt: time.Now().Add(-10 * time.Hour),
				Counts:    100,
			},
		}

		mockService.EXPECT().
			GetRecordsByDateAndCount(recordRequest).
			Return(expectedRecords, nil).Times(1)
		jsonBytes, _ := json.Marshal(recordRequest)
		body := new(bytes.Buffer)
		body.Write(jsonBytes)
		rec, req := createHTTPReqForTest(http.MethodPost, "/record", body)
		handler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("error", func(t *testing.T) {
		recordRequest := record.RecordRequest{
			StartDate: "2021-05-05",
			EndDate:   "2021-05-06",
			MinCount:  10,
			MaxCount:  20,
		}

		mockService.EXPECT().
			GetRecordsByDateAndCount(recordRequest).
			Return(nil, fmt.Errorf("test")).Times(1)
		jsonBytes, _ := json.Marshal(recordRequest)
		body := new(bytes.Buffer)
		body.Write(jsonBytes)
		rec, req := createHTTPReqForTest(http.MethodPost, "/record", body)
		handler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
func createHTTPReqForTest(method, endpoint string, body io.Reader) (w *httptest.ResponseRecorder, req *http.Request) {
	if body == nil {
		const size = 512
		body = bytes.NewBuffer(make([]byte, size))
	}
	req = httptest.NewRequest(method, endpoint, body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	rec := httptest.NewRecorder()
	return rec, req
}

func createMockService(t *testing.T) *mocks.MockRecordService {
	c := gomock.NewController(t)
	service := mocks.NewMockRecordService(c)
	return service
}
