package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-live-chat/internal/configs"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/infraestructure/wrappers"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

type MockCursor struct {
	mock.Mock
}

func (m *MockCursor) All(ctx context.Context, results interface{}) error {
	args := m.Called(ctx, results)
	return args.Error(0)
}

type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) Disconnect(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockMongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	args := m.Called(ctx, rp)
	return args.Error(0)
}

func (m *MockMongoClient) Database(name string, opts ...*options.DatabaseOptions) wrappers.MongoDatabaseInterface {
	args := m.Called(name)
	return args.Get(0).(wrappers.MongoDatabaseInterface)
}

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Collection(name string, opts ...*options.CollectionOptions) wrappers.MongoCollectionInterface {
	args := m.Called(name)
	return args.Get(0).(wrappers.MongoCollectionInterface)
}

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (wrappers.MongoCursorInterface, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(wrappers.MongoCursorInterface), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func TestChatroomRepository_Create(t *testing.T) {
	// Arrange
	mockClient := new(MockMongoClient)
	mockDatabase := new(MockDatabase)
	mockCollection := new(MockCollection)

	dbConnections := databases.MongoDBConnections{
		OpenChat: mockClient,
	}

	repo := NewChatroomRepository(&dbConnections, &configs.Config{
		OpenChatMongoDB: &configs.MongoDBConfig{
			Database: "chatDB",
		},
	})

	chatroom := model.Chatroom{
		Name:        "Test Room",
		Description: "Test description",
		Owner:       "user-id",
	}

	// Mock the InsertOne behavior
	mockClient.On("Database", "chatDB").Return(mockDatabase)
	mockDatabase.On("Collection", "chatrooms").Return(mockCollection)
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil)

	// Act
	createdChatroom, err := repo.Create(chatroom, context.Background())

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, createdChatroom)
	assert.Equal(t, "Test Room", createdChatroom.Name)
	mockClient.AssertExpectations(t)
	mockDatabase.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}

func TestChatroomRepository_GetByFilter(t *testing.T) {
	// Arrange
	mockClient := new(MockMongoClient)
	mockDatabase := new(MockDatabase)
	mockCollection := new(MockCollection)
	mockCursor := new(MockCursor)

	dbConnections := databases.MongoDBConnections{
		OpenChat: mockClient,
	}

	repo := NewChatroomRepository(&dbConnections, &configs.Config{
		OpenChatMongoDB: &configs.MongoDBConfig{
			Database: "chatDB",
		},
	})

	chatrooms := []model.Chatroom{
		{Id: primitive.NewObjectID(), Name: "Room 1", Description: "Room 1 description", Owner: "owner1"},
		{Id: primitive.NewObjectID(), Name: "Room 2", Description: "Room 2 description", Owner: "owner2"},
	}

	// Mock Find behavior
	mockClient.On("Database", "chatDB").Return(mockDatabase)
	mockDatabase.On("Collection", "chatrooms").Return(mockCollection)
	mockCollection.On("Find", mock.Anything, mock.Anything).Return(mockCursor, nil)
	mockCursor.On("All", mock.Anything, mock.Anything).Return(nil, nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*[]model.Chatroom)
		*arg = chatrooms
	})

	// Act
	rooms, err := repo.GetByFilter(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Len(t, rooms, 2)
	mockClient.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}

func TestChatroomRepository_GetById(t *testing.T) {
	// Arrange
	mockClient := new(MockMongoClient)
	mockDatabase := new(MockDatabase)
	mockCollection := new(MockCollection)

	dbConnections := databases.MongoDBConnections{
		OpenChat: mockClient,
	}

	repo := NewChatroomRepository(&dbConnections, &configs.Config{
		OpenChatMongoDB: &configs.MongoDBConfig{
			Database: "chatDB",
		},
	})

	searchId := primitive.NewObjectID()
	chatroom := model.Chatroom{
		Id:          searchId,
		Name:        "Test Room",
		Description: "Test description",
		Owner:       "user-id",
	}

	expectedSearch := bson.M{
		"_id": searchId,
	}

	// Mock the InsertOne behavior
	mockClient.On("Database", "chatDB").Return(mockDatabase)
	mockDatabase.On("Collection", "chatrooms").Return(mockCollection)
	mockCollection.On("FindOne", mock.Anything, expectedSearch).Return(mongo.NewSingleResultFromDocument(chatroom, nil, nil))

	// Act
	chatroomResp, err := repo.GetById(searchId.Hex(), context.Background())

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, chatroomResp)
	mockClient.AssertExpectations(t)
	mockDatabase.AssertExpectations(t)
	mockCollection.AssertExpectations(t)

}

func TestChatroomRepository_Update(t *testing.T) {
	// Arrange
	mockClient := new(MockMongoClient)
	mockDatabase := new(MockDatabase)
	mockCollection := new(MockCollection)

	dbConnections := databases.MongoDBConnections{
		OpenChat: mockClient,
	}

	repo := NewChatroomRepository(&dbConnections, &configs.Config{
		OpenChatMongoDB: &configs.MongoDBConfig{
			Database: "chatDB",
		},
	})

	chatroom := model.Chatroom{
		Id:          primitive.NewObjectID(),
		Name:        "Test Room",
		Description: "Updated description",
		Owner:       "owner-id",
	}

	// Mock Update behavior
	mockClient.On("Database", "chatDB").Return(mockDatabase)
	mockDatabase.On("Collection", "chatrooms").Return(mockCollection)
	mockCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil)

	// Act
	updatedChatroom, err := repo.Update(chatroom, context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Test Room", updatedChatroom.Name)
	mockClient.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}
