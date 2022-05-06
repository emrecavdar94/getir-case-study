package inmemory_test

import (
	"fmt"
	inmemory "getir-assignment/internal/in_memory"
	"getir-assignment/internal/in_memory/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestService_Get_GivenKeyValue_ThenReturnKeyValueData(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mocks.NewMockRepository(mockController)

	service := inmemory.NewService(mockRepository, zap.NewNop())
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	mockRepository.EXPECT().
		Get(expectedResult.Key).
		Return(&expectedResult, nil).Times(1)

	actualResult, err := service.Get("test")
	assert.Nil(t, err)
	assert.Equal(t, &expectedResult, actualResult)
}

func TestService_Get_GivenNotExistKey_ThenReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mocks.NewMockRepository(mockController)
	service := inmemory.NewService(mockRepository, zap.NewNop())
	mockRepository.EXPECT().
		Get("test").
		Return(nil, fmt.Errorf("test")).Times(1)
	_, err := service.Get("test")
	assert.NotNil(t, err)
}

func TestService_Set_GivenKeyValue_ThenReturnKeyValueData(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mocks.NewMockRepository(mockController)
	service := inmemory.NewService(mockRepository, zap.NewNop())
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	mockRepository.EXPECT().
		Set(expectedResult.Key, expectedResult.Value).
		Return(&expectedResult, nil).Times(1)
	actualResult, err := service.Set("test", "value")
	assert.Nil(t, err)
	assert.Equal(t, &expectedResult, actualResult)
}
func TestService_Set_GivenAlreadyExistKey_ThenReturnError(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mocks.NewMockRepository(mockController)
	service := inmemory.NewService(mockRepository, zap.NewNop())
	mockRepository.EXPECT().
		Set("test", "value").
		Return(nil, fmt.Errorf("test")).Times(1)
	_, err := service.Set("test", "value")
	assert.NotNil(t, err)
}
