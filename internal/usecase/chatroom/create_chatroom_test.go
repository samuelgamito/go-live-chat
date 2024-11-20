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

type MockChatroomRepository struct {
	mock.Mock
}

func (m *MockChatroomRepository) Create(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error) {
	args := m.Called(chatroom, ctx)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Chatroom), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateChatRoomUseCase_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepository)
	useCase := NewCreateChatroomUseCase(mockRepo)

	chatroom := model.Chatroom{
		Name:        "Test Room",
		Description: "A test chatroom",
		Owner:       "owner-id",
	}

	expectedResponse := &model.Chatroom{
		Id:          primitive.NewObjectID(),
		Name:        chatroom.Name,
		Description: chatroom.Description,
		Owner:       chatroom.Owner,
		Members: []model.Member{
			{Id: chatroom.Owner, SinceAt: time.Now()},
		},
	}

	mockRepo.On("Create", mock.AnythingOfType("model.Chatroom"), mock.Anything).Return(expectedResponse, nil)

	// Act
	id, err := useCase.Execute(chatroom, context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Id.Hex(), id)
	mockRepo.AssertExpectations(t)
}

func TestCreateChatRoomUseCase_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepository)
	useCase := NewCreateChatroomUseCase(mockRepo)

	chatroom := model.Chatroom{
		Name:        "Test Room",
		Description: "A test chatroom",
		Owner:       "owner-id",
	}

	mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(nil, errors.New("repository error"))

	// Act
	id, err := useCase.Execute(chatroom, context.Background())

	// Assert
	assert.Empty(t, id, "Expected ID to be empty on error")
	assert.NotNil(t, err, "Expected an error to be returned")
	assert.Equal(t, []string{"Error creating chatroom"}, err.Messages)
	mockRepo.AssertExpectations(t)
}

func TestCreateChatRoomUseCase_AppendOwnerToMembers(t *testing.T) {
	// Arrange
	mockRepo := new(MockChatroomRepository)
	useCase := NewCreateChatroomUseCase(mockRepo)

	chatroom := model.Chatroom{
		Name:        "Test Room",
		Description: "A test chatroom",
		Owner:       "owner-id",
	}

	expectedChatroom := chatroom
	expectedChatroom.Members = append(expectedChatroom.Members, model.Member{
		Id:      chatroom.Owner,
		SinceAt: time.Now(),
	})

	mockRepo.
		On("Create", mock.AnythingOfType("model.Chatroom"), mock.Anything).
		Return(&model.Chatroom{Id: primitive.NewObjectID()}, nil)

	// Act
	_, _ = useCase.Execute(chatroom, context.Background())

	// Assert that the owner was appended to the members list
	mockRepo.AssertCalled(t, "Create", mock.MatchedBy(func(c model.Chatroom) bool {
		return len(c.Members) == 1 && c.Members[0].Id == chatroom.Owner
	}), mock.Anything)
}
