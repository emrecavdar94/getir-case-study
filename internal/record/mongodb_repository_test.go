package record_test

import (
	"context"
	"getir-assignment/config"
	"getir-assignment/internal/record"
	"log"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBTestSuite struct {
	suite.Suite
	container  testcontainers.Container
	conf       config.MongoConfig
	client     *mongo.Client
	repository record.Repository
}

func (s *MongoDBTestSuite) SetupSuite() {
	SetupTestMongoDBInstance(s.T())
	s.conf = s.prepareConfig()
	s.client = s.createMongoDBClient()
	s.repository = record.NewRepository(s.client, s.conf)
}

func (s *MongoDBTestSuite) prepareConfig() config.MongoConfig {
	return config.MongoConfig{
		URI:        "mongodb://localhost:27017",
		DBName:     "getir-case-study",
		Collection: "records",
	}
}

func (s *MongoDBTestSuite) createMongoDBClient() *mongo.Client {
	//	credential := options.Credential{Username: config.Username, Password: config.Password}
	clientOptions := options.Client().ApplyURI(s.conf.URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SetupTestMongoDBInstance(t *testing.T) {
	startTestContainer(t)
}

func startTestContainer(t *testing.T) testcontainers.Container {
	port, err := nat.NewPort("tcp", "27017")
	assert.Nil(t, err)

	req := testcontainers.ContainerRequest{
		Image: "mongo",
		ExposedPorts: []string{
			"27017:27017/tcp",
		},
		WaitingFor: wait.ForListeningPort(port),
	}

	ctx := context.Background()
	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})

	assert.Nil(t, err)
	return container
}
func TestMongoDB(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(MongoDBTestSuite))
}
