package usecase_chatroom

import (
	"context"
	"errors"
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChatroomRepositorySearch struct {
	mock.Mock
}

func (m *MockChatroomRepositorySearch) GetById(id string, ctx context.Context) (*model.Chatroom, error) {
	args := m.Called(id, ctx)
	return args.Get(0).(*model.Chatroom), args.Error(1)
}

func (m *MockChatroomRepositorySearch) GetByFilter(ctx context.Context) ([]model.Chatroom, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Chatroom), args.Error(1)
}

func TestRetrieveChatroom_ExecuteById_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	retrieveChatroom := NewRetrieveChatroom(mockRepo)

	chatroom := &model.Chatroom{
		Id:          primitive.NewObjectID(),
		Name:        "Test Chatroom",
		Description: "Test Description",
		Owner:       "owner-id",
	}

	mockRepo.On("GetById", "test-id", mock.Anything).Return(chatroom, nil)

	// Act
	result, err := retrieveChatroom.ExecuteById("test-id", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, chatroom, result)
	mockRepo.AssertExpectations(t)
}

func TestRetrieveChatroom_ExecuteById_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	retrieveChatroom := NewRetrieveChatroom(mockRepo)

	mockRepo.On("GetById", "test-id", mock.Anything).
		Return((*model.Chatroom)(nil), errors.New("repository error"))

	// Act
	result, err := retrieveChatroom.ExecuteById("test-id", context.Background())

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, misc.DefaultError(), err)
	mockRepo.AssertExpectations(t)
}

func TestRetrieveChatroom_ExecuteByFilter_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	retrieveChatroom := NewRetrieveChatroom(mockRepo)

	chatrooms := []model.Chatroom{
		{Id: primitive.NewObjectID(), Name: "Chatroom 1", Description: "Description 1", Owner: "Owner 1"},
		{Id: primitive.NewObjectID(), Name: "Chatroom 2", Description: "Description 2", Owner: "Owner 2"},
	}

	mockRepo.On("GetByFilter", mock.Anything).Return(chatrooms, nil)

	// Act
	result, err := retrieveChatroom.ExecuteByFilter(context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, chatrooms, result)
	mockRepo.AssertExpectations(t)
}

func TestRetrieveChatroom_ExecuteByFilter_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositorySearch)
	retrieveChatroom := NewRetrieveChatroom(mockRepo)

	mockRepo.On("GetByFilter", mock.Anything).
		Return([]model.Chatroom{}, errors.New("repository error"))

	// Act
	result, err := retrieveChatroom.ExecuteByFilter(context.Background())

	// Assert
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, misc.DefaultError(), err)
	mockRepo.AssertExpectations(t)
}
