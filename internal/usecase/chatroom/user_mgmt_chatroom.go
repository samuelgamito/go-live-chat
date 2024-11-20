package usecase_chatroom

import (
	"context"
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
	"time"
)

type UserManagementChatroomUseCase struct {
	repo usecase.ChatroomRepository
}

func NewUserManagementChatroomUseCase(repo usecase.ChatroomRepository) *UserManagementChatroomUseCase {
	return &UserManagementChatroomUseCase{
		repo: repo,
	}
}

func (u *UserManagementChatroomUseCase) Join(roomId string, userId string, ctx context.Context) *model.Error {

	chatroom, err := u.repo.GetById(roomId, ctx)

	if err != nil {
		return misc.DefaultError()
	}

	members := append(chatroom.Members, model.Member{
		Id:      userId,
		SinceAt: time.Now(),
	})

	chatroom.Members = members

	_, err = u.repo.Update(*chatroom, ctx)

	if err != nil {
		return misc.DefaultError()
	}
	return nil
}

func (u *UserManagementChatroomUseCase) Leave(roomId string, userId string, ctx context.Context) *model.Error {

	return nil
}
