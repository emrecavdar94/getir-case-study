package record

import (
	"context"
	"fmt"
	"getir-assignment/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repository struct {
	client *mongo.Client
	config config.MongoConfig
}

func NewRepository(client *mongo.Client, config config.MongoConfig) Repository {
	return &repository{
		client: client,
		config: config,
	}
}

func (r *repository) GetRecordsByDateAndCount(startDate, endDate time.Time, minCount, maxCount int) (records []*Record, err error) {
	coll := r.client.Database(r.config.DBName).Collection(r.config.Collection)
	records = []*Record{}
	pipe := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gte": minCount,
					"$lte": maxCount,
				},
			},
		},
	}

	fmt.Println(startDate, endDate)
	cursor, err := coll.Aggregate(context.TODO(), pipe)
	defer cursor.Close(context.TODO())
	if err != nil {
		return records, err
	}

	if err = cursor.All(context.TODO(), &records); err != nil {
		return nil, err
	}

	return records, nil
}
