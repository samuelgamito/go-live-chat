package usecase_chatroom

import (
	"context"
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
	"time"
)

type UserManagementChatroomUseCase struct {
	repo usecase.ChatroomRepositoryUpdate
}

func NewUserManagementChatroomUseCase(repo usecase.ChatroomRepositoryUpdate) *UserManagementChatroomUseCase {
	return &UserManagementChatroomUseCase{
		repo: repo,
	}
}

func (u *UserManagementChatroomUseCase) Join(roomId string, userId string, ctx context.Context) (*model.Chatroom, *model.Error) {

	chatroom, err := u.repo.GetById(roomId, ctx)

	if err != nil {
		return nil, misc.DefaultError()
	}

	if misc.SliceContainsString(chatroom.Members, "Id", userId) > 0 {
		return chatroom, nil
	}

	members := append(chatroom.Members, model.Member{
		Id:      userId,
		SinceAt: time.Now(),
	})

	chatroom.Members = members

	chatroom, err = u.repo.Update(*chatroom, ctx)

	if err != nil {
		return chatroom, misc.DefaultError()
	}
	return chatroom, nil
}

func (u *UserManagementChatroomUseCase) Leave(roomId string, userId string, ctx context.Context) (*model.Chatroom, *model.Error) {

	chatroom, err := u.repo.GetById(roomId, ctx)

	if err != nil {
		return chatroom, misc.DefaultError()
	}

	if len(chatroom.Members) == 0 {
		return chatroom, nil
	}

	if pos := misc.SliceContainsString(chatroom.Members, "Id", userId); pos > 0 {
		chatroom.Members = misc.RemoveByIndex(chatroom.Members, pos).([]model.Member)
		update, err := u.repo.Update(*chatroom, ctx)
		if err != nil {
			return nil, misc.DefaultError()
		}
		return update, nil
	}

	return chatroom, nil
}
