package usecase_chatroom

import (
	"context"
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
)

type RetrieveChatroom struct {
	repo usecase.ChatroomRepository
}

func NewRetrieveChatroom(repo usecase.ChatroomRepository) *RetrieveChatroom {
	return &RetrieveChatroom{
		repo: repo,
	}
}

func (r *RetrieveChatroom) ExecuteById(id string, ctx context.Context) (*model.Chatroom, *model.Error) {

	d, err := r.repo.GetById(id, ctx)

	if err != nil {
		return nil, misc.DefaultError()
	}

	return d, nil
}

func (r *RetrieveChatroom) ExecuteByFilter(ctx context.Context) ([]model.Chatroom, *model.Error) {

	d, err := r.repo.GetByFilter(ctx)

	if err != nil {
		return nil, misc.DefaultError()
	}

	return d, nil
}
