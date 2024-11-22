package usecase_chatroom

import (
	"context"
	"errors"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChatroomRepositoryUpdate struct {
	mock.Mock
}

func (m *MockChatroomRepositoryUpdate) GetById(id string, ctx context.Context) (*model.Chatroom, error) {
	args := m.Called(id, ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Chatroom), nil
}

func (m *MockChatroomRepositoryUpdate) Update(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error) {
	args := m.Called(chatroom, ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Chatroom), nil
}

func TestUserManagementChatroomUseCase_Join_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id: primitive.NewObjectID(),
		Members: []model.Member{
			{Id: "existing-user", SinceAt: time.Now()},
		},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)
	mockRepo.On("Update", mock.AnythingOfType("model.Chatroom"), mock.Anything).Return(chatroom, nil)

	// Act
	chatroomResp, err := useCase.Join("room1", "new-user", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Len(t, chatroom.Members, 2)
	assert.Equal(t, "new-user", chatroomResp.Members[1].Id)
	assert.Equal(t, chatroom.Id, chatroomResp.Id)
	mockRepo.AssertExpectations(t)
}

func TestUserManagementChatroomUseCase_Join_Success_AlreadyJoined(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id: primitive.NewObjectID(),
		Members: []model.Member{
			{Id: "existing-user", SinceAt: time.Now()},
			{Id: "new-user", SinceAt: time.Now()},
		},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)

	// Act
	chatroomResp, err := useCase.Join("room1", "new-user", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Len(t, chatroom.Members, 2)
	assert.Equal(t, "new-user", chatroomResp.Members[1].Id)
	assert.Equal(t, chatroom.Id, chatroomResp.Id)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Update", mock.AnythingOfType("model.Chatroom"), mock.Anything)
}

func TestUserManagementChatroomUseCase_Join_GetByIdError(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	mockRepo.On("GetById", "room1", mock.Anything).
		Return(nil, errors.New("not found"))

	// Act
	chatroomResp, err := useCase.Join("room1", "new-user", context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, chatroomResp)
	mockRepo.AssertExpectations(t)
}

func TestUserManagementChatroomUseCase_Join_UpdateError(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id: primitive.NewObjectID(),
		Members: []model.Member{
			{Id: "existing-user", SinceAt: time.Now()},
		},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)
	mockRepo.On("Update", mock.AnythingOfType("model.Chatroom"), mock.Anything).
		Return(nil, errors.New("update failed"))

	// Act
	chatroomResp, err := useCase.Join("room1", "new-user", context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, chatroomResp)
	mockRepo.AssertExpectations(t)
}

func TestUserManagementChatroomUseCase_Leave_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id: primitive.NewObjectID(),
		Members: []model.Member{
			{Id: "existing-user", SinceAt: time.Now()},
			{Id: "new-user", SinceAt: time.Now()},
		},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)
	mockRepo.On("Update", mock.AnythingOfType("model.Chatroom"), mock.Anything).Return(chatroom, nil)

	// Act
	chatroomResp, err := useCase.Leave("room1", "new-user", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Len(t, chatroom.Members, 1)
	assert.Equal(t, "existing-user", chatroomResp.Members[0].Id)
	assert.Equal(t, chatroom.Id, chatroomResp.Id)
	mockRepo.AssertExpectations(t)
}

func TestUserManagementChatroomUseCase_Leave_Success_AlreadyLeaved(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id: primitive.NewObjectID(),
		Members: []model.Member{
			{Id: "existing-user", SinceAt: time.Now()},
		},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)

	// Act
	chatroomResp, err := useCase.Leave("room1", "new-user", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Len(t, chatroom.Members, 1)
	assert.Equal(t, "existing-user", chatroomResp.Members[0].Id)
	assert.Equal(t, chatroom.Id, chatroomResp.Id)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Update", mock.AnythingOfType("model.Chatroom"), mock.Anything)
}

func TestUserManagementChatroomUseCase_Leave_Success_EmptyMemberList(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id:      primitive.NewObjectID(),
		Members: []model.Member{},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)

	// Act
	chatroomResp, err := useCase.Leave("room1", "new-user", context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Len(t, chatroom.Members, 0)
	assert.Equal(t, chatroom.Id, chatroomResp.Id)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Update", mock.AnythingOfType("model.Chatroom"), mock.Anything)
}

func TestUserManagementChatroomUseCase_Leave_GetByIdError(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	mockRepo.On("GetById", "room1", mock.Anything).
		Return(nil, errors.New("not found"))

	// Act
	chatroomResp, err := useCase.Leave("room1", "new-user", context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, chatroomResp)
	mockRepo.AssertExpectations(t)
}

func TestUserManagementChatroomUseCase_Leave_UpdateError(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepositoryUpdate)
	useCase := NewUserManagementChatroomUseCase(mockRepo)

	chatroom := &model.Chatroom{
		Id: primitive.NewObjectID(),
		Members: []model.Member{
			{Id: "existing-user", SinceAt: time.Now()},
			{Id: "new-user", SinceAt: time.Now()},
		},
	}

	mockRepo.On("GetById", "room1", mock.Anything).Return(chatroom, nil)
	mockRepo.On("Update", mock.AnythingOfType("model.Chatroom"), mock.Anything).
		Return(nil, errors.New("update failed"))

	// Act
	chatroomResp, err := useCase.Leave("room1", "new-user", context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, chatroomResp)
	mockRepo.AssertExpectations(t)
}
