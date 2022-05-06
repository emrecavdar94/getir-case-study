package record_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MongoRecordData struct {
	Key       string    `json:"key" bson:"key"`
	Value     string    `json:"value" bson:"value"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Counts    []int     `json:"counts" bson:"counts"`
}

func (s *MongoDBTestSuite) TestGetRecords() {
	s.T().Run("given records then should return all of it ", func(t *testing.T) {
		startDate := time.Now().Add(-24 * time.Hour)
		endDate := time.Now()
		minCount := 100
		maxCount := 500

		records := []*MongoRecordData{
			{
				Key:       "111",
				Value:     "1112",
				CreatedAt: time.Now().Add(-12 * time.Hour),
				Counts: []int{
					100, 100,
				},
			},
		}
		coll := s.client.Database(s.conf.DBName).Collection(s.conf.Collection)
		docs := make([]interface{}, len(records))
		for i, v := range records {
			docs[i] = v
		}
		_, err := coll.InsertMany(context.Background(), docs)
		assert.Nil(t, err)

		actualRecords, err := s.repository.GetRecordsByDateAndCount(startDate, endDate, minCount, maxCount)
		assert.Nil(t, err)
		assert.Equal(t, len(records), len(actualRecords))
	})
}
