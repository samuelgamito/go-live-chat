package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-live-chat/internal/handlers/dto"
	"go-live-chat/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for CreateChatRoomUseCase
type ChatRoomUseCase struct {
	mock.Mock
}

func (m *ChatRoomUseCase) Execute(model model.Chatroom, ctx context.Context) (string, *model.Error) {
	args := m.Called(model, ctx)
	return args.String(0), nil
}

func TestCreateChatroom_Success(t *testing.T) {
	// Arrange
	mockUseCase := new(ChatRoomUseCase)
	handler := NewChatRoomHandler(mockUseCase)
	router := http.NewServeMux()
	router.HandleFunc("/api/chatrooms", handler.createChatroom)

	createRequest := dto.CreateChatRoomRequestDTO{
		Name:        "Test Room",
		Owner:       "sgamito",
		Description: "adasdasd",
	}

	mockUseCase.On("Execute", mock.Anything, mock.Anything).Return("12345", nil)

	body, _ := json.Marshal(createRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/chatrooms", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	assert.Equal(t, http.StatusCreated, recorder.Code)
	var resp dto.CreatedChatRoomDTO
	err := json.NewDecoder(recorder.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "12345", resp.Id)
	mockUseCase.AssertExpectations(t)
}

func TestCreateChatroom_InvalidRequest(t *testing.T) {
	// Arrange
	mockUseCase := new(ChatRoomUseCase)
	handler := NewChatRoomHandler(mockUseCase)
	router := http.NewServeMux()
	router.HandleFunc("/api/chatrooms", handler.createChatroom)

	req := httptest.NewRequest(http.MethodPost, "/api/chatrooms", bytes.NewReader([]byte("invalid JSON")))
	recorder := httptest.NewRecorder()

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Invalid request payload")
}

func TestCreateChatroom_ValidationFails(t *testing.T) {
	// Arrange
	mockUseCase := new(ChatRoomUseCase)
	handler := NewChatRoomHandler(mockUseCase)
	router := http.NewServeMux()
	router.HandleFunc("/api/chatrooms", handler.createChatroom)

	// Create an invalid request
	createRequest := dto.CreateChatRoomRequestDTO{} // missing required fields for validation to fail
	body, _ := json.Marshal(createRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/chatrooms", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, recorder.Code) // assuming validation errors return 422
}

func TestCreateChatroom_UseCaseError(t *testing.T) {
	// Arrange
	mockUseCase := new(ChatRoomUseCase)
	handler := NewChatRoomHandler(mockUseCase)
	router := http.NewServeMux()
	router.HandleFunc("/api/chatrooms", handler.createChatroom)

	createRequest := dto.CreateChatRoomRequestDTO{
		Name: "Test Room",
	}
	mockUseCase.On("Execute", createRequest.ToModel(), mock.Anything).Return("", errors.New("use case error"))

	body, _ := json.Marshal(createRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/chatrooms", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	assert.Equal(t, "{\"messages\":[\"Field Owner is required.\"]}", recorder.Body.String())

}
