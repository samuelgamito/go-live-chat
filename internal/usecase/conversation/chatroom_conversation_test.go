package usecase_conversation

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type MockChatroomRepositorySearch struct {
	mock.Mock
}

func (m *MockChatroomRepositorySearch) GetById(id string, ctx context.Context) (*model.Chatroom, error) {
	args := m.Called(id, ctx)
	if (args.Get(0)) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Chatroom), nil
}

func (m *MockChatroomRepositorySearch) GetByFilter(ctx context.Context) ([]model.Chatroom, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Chatroom), args.Error(1)
}

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return nil
}

func (m *MockRedisClient) Process(ctx context.Context, cmd redis.Cmder) error {
	return nil
}

func (m *MockRedisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return nil
}

func (m *MockRedisClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	args := m.Called(ctx, channel, message)
	return args.Get(0).(*redis.IntCmd)
}

type MockConversationRepository struct {
	mock.Mock
}

func (m *MockConversationRepository) SaveMessageToUser(message model.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockConversationRepository) BatchSaveMessage(message []model.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

func TestFindMembers_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	rdb := &databases.RedisClient{}

	conversationUseCase := NewChatroomConversationUseCase(mockRepo, nil, rdb)

	chatrooms := model.Chatroom{
		Id:          primitive.NewObjectID(),
		Name:        "Chatroom 1",
		Description: "Description 1",
		Owner:       "Owner 1",
		Members: []model.Member{
			{Id: "member1"},
			{Id: "member1"},
		},
	}

	mockRepo.On("GetById", "chatroom", mock.Anything).Return(&chatrooms, nil)

	// Act
	members, err := conversationUseCase.FindMembers("chatroom", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, chatrooms.Members[0].Id, members[0].Id)
	assert.Equal(t, chatrooms.Members[1].Id, members[1].Id)
	mockRepo.AssertExpectations(t)

}

func TestFindMembers_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	rdb := &databases.RedisClient{
		NotifyClientsRedis: new(MockRedisClient),
	}

	conversationUseCase := NewChatroomConversationUseCase(mockRepo, nil, rdb)

	mockRepo.On("GetById", "chatroom", mock.Anything).Return(nil, errors.New("error"))

	// Act
	members, err := conversationUseCase.FindMembers("chatroom", context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, members)
	mockRepo.AssertExpectations(t)

}

func TestPrepareMessage(t *testing.T) {
	// Arrange
	rdb := &databases.RedisClient{}
	members := []model.Member{
		{Id: "member1"},
		{Id: "member1"},
	}
	conversationUseCase := NewChatroomConversationUseCase(nil, nil, rdb)

	// Act
	messages := conversationUseCase.PrepareMessage(members, "message", "from")

	// Assert

	assert.Equal(t, len(messages), len(members))
	assert.Equal(t, messages[0].To, members[0].Id)
	assert.Equal(t, messages[0].Content, "message")
	assert.Equal(t, messages[0].From, "from")
	assert.Equal(t, messages[1].To, members[1].Id)
	assert.Equal(t, messages[1].Content, "message")
	assert.Equal(t, messages[1].From, "from")

}

func TestPublishMessage_Success(t *testing.T) {

	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	mockRedis := new(MockRedisClient)

	rdb := &databases.RedisClient{
		NotifyClientsRedis: mockRedis,
	}

	redisResp := &redis.IntCmd{}
	redisResp.SetErr(nil)

	messages := []model.Message{
		{Id: primitive.NewObjectID(), From: "from", To: "to1", Content: "message", Type: "chatroom"},
		{Id: primitive.NewObjectID(), From: "from", To: "to2", Content: "message", Type: "chatroom"},
	}

	mockRedis.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(redisResp)

	conversationUseCase := NewChatroomConversationUseCase(mockRepo, nil, rdb)

	// Act
	err := conversationUseCase.PublishMessage(messages, context.Background())

	// Assert
	assert.Nil(t, err)
	mockRedis.AssertExpectations(t)
	mockRedis.AssertNumberOfCalls(t, "Publish", 2)
}

func TestPublishMessage_RedisErr(t *testing.T) {

	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	mockRedis := new(MockRedisClient)

	rdb := &databases.RedisClient{
		NotifyClientsRedis: mockRedis,
	}

	redisResp := &redis.IntCmd{}
	redisResp.SetErr(errors.New("error"))

	messages := []model.Message{
		{Id: primitive.NewObjectID(), From: "from", To: "to1", Content: "message", Type: "chatroom"},
		{Id: primitive.NewObjectID(), From: "from", To: "to2", Content: "message", Type: "chatroom"},
	}

	mockRedis.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(redisResp)

	conversationUseCase := NewChatroomConversationUseCase(mockRepo, nil, rdb)

	// Act
	err := conversationUseCase.PublishMessage(messages, context.Background())

	// Assert
	assert.NotNil(t, err)
	mockRedis.AssertExpectations(t)
	mockRedis.AssertNumberOfCalls(t, "Publish", 1)
}
