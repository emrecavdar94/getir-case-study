package inmemory_test

import (
	inmemory "getir-assignment/internal/in_memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Get_GivenKeyValue_ThenReturnKeyValueData(t *testing.T) {
	repository := inmemory.NewRepository()
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	_, _ = repository.Set("test", "value")
	actualResult, err := repository.Get("test")
	assert.Nil(t, err)
	assert.Equal(t, &expectedResult, actualResult)
}
func TestRepository_Get_GivenNotExistKey_ThenReturnError(t *testing.T) {
	repository := inmemory.NewRepository()
	_, _ = repository.Set("test", "value")
	_, err := repository.Get("test2")
	assert.NotNil(t, err)
}

func TestRepository_Set_GivenKeyValue_ThenReturnKeyValueData(t *testing.T) {
	repository := inmemory.NewRepository()
	expectedResult := inmemory.InMemory{
		Key:   "test",
		Value: "value",
	}
	actualResult, err := repository.Set("test", "value")
	assert.Nil(t, err)
	assert.Equal(t, &expectedResult, actualResult)
}
func TestRepository_Set_GivenAlreadyExistKey_ThenReturnError(t *testing.T) {
	repository := inmemory.NewRepository()
	_, _ = repository.Set("test", "value")
	_, err := repository.Set("test", "value2")
	assert.NotNil(t, err)
}
