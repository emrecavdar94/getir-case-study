package inmemory_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	inmemory "getir-assignment/internal/in_memory"
	"getir-assignment/internal/in_memory/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHandler_Get_GivenKeyValue_ThenReturnKeyValueData(t *testing.T) {
	mockService := createMockService(t)

	handler := inmemory.NewHandler(mockService, zap.NewNop())
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	mockService.EXPECT().
		Get(expectedResult.Key).
		Return(&expectedResult, nil).Times(1)

	rec, req := createHTTPReqForTest(http.MethodGet, "/in-memory?key=test", http.NoBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestHandler_Get_GivenNotExistKey_ThenReturnError(t *testing.T) {
	mockService := createMockService(t)

	handler := inmemory.NewHandler(mockService, zap.NewNop())
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	mockService.EXPECT().
		Get(expectedResult.Key).
		Return(nil, fmt.Errorf("test")).Times(1)

	rec, req := createHTTPReqForTest(http.MethodGet, "/in-memory?key=test", http.NoBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestHandler_Set_GivenKeyValue_ThenReturnKeyValueData(t *testing.T) {
	mockService := createMockService(t)

	handler := inmemory.NewHandler(mockService, zap.NewNop())
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	mockService.EXPECT().
		Set(expectedResult.Key, expectedResult.Value).
		Return(&expectedResult, nil).Times(1)
	jsonBytes, _ := json.Marshal(expectedResult)
	body := new(bytes.Buffer)
	body.Write(jsonBytes)
	rec, req := createHTTPReqForTest(http.MethodPost, "/in-memory", body)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestHandler_Set_GivenAlreadyExistKey_ThenReturnError(t *testing.T) {
	mockService := createMockService(t)

	handler := inmemory.NewHandler(mockService, zap.NewNop())
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	mockService.EXPECT().
		Set(expectedResult.Key, expectedResult.Value).
		Return(nil, fmt.Errorf("test")).Times(1)
	jsonBytes, _ := json.Marshal(expectedResult)
	body := new(bytes.Buffer)
	body.Write(jsonBytes)
	rec, req := createHTTPReqForTest(http.MethodPost, "/in-memory", body)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
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

func createMockService(t *testing.T) *mocks.MockMemoryService {
	c := gomock.NewController(t)
	service := mocks.NewMockMemoryService(c)
	return service
}
