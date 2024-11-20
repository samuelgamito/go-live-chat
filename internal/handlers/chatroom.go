package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go-live-chat/internal/constants"
	"go-live-chat/internal/handlers/dto"
	"go-live-chat/internal/misc"
	"go.uber.org/fx"
	"net/http"
)

type ChatRoomHandler struct {
	createChatroomUseCase         CreateChatroomUseCase
	retrieveChatroomUseCase       RetrieveChatroomUseCase
	userManagementChatroomUseCase UserManagementChatroomUseCase
}

func NewChatRoomHandler(
	createChatroomUseCase CreateChatroomUseCase,
	retrieveChatroomUseCase RetrieveChatroomUseCase,
	userManagementChatroomUseCase UserManagementChatroomUseCase,
) *ChatRoomHandler {
	return &ChatRoomHandler{
		createChatroomUseCase:         createChatroomUseCase,
		retrieveChatroomUseCase:       retrieveChatroomUseCase,
		userManagementChatroomUseCase: userManagementChatroomUseCase,
	}
}

func registerChatRoomHandlers(c *ChatRoomHandler, h *Handler) {
	h.Runner.Group(func(r chi.Router) {
		//r.Use(tokenValidationMiddleware)
		r.Route("/api/chatrooms", func(r chi.Router) {
			r.Post("/", c.createChatroom)
			r.Get("/", c.listChatroom)
			r.Route("/{roomId}", func(r chi.Router) {
				r.Use(chatroomCtx)

				r.Get("/", c.chatRoomDetails)
				r.Post("/leave", c.leaveChatroom)
				r.Post("/join", c.joinChatroom)
			})
		})
	})
}

func (c *ChatRoomHandler) createChatroom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var createRequest dto.CreateChatRoomRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if errValid := createRequest.IsValid(); errValid != nil {
		misc.WriteJSONResponse(w, errValid.StatusCode, errValid.Body)
		return
	}

	id, errUseCase := c.createChatroomUseCase.Execute(createRequest.ToModel(), ctx)
	if errUseCase != nil {
		errResp := dto.ErrorResponse{}
		errResp.FromModel(errUseCase)
		misc.WriteJSONResponse(w, errResp.StatusCode, errResp.Body)
		return
	}

	resp := dto.CreatedChatRoomDTO{
		Id: id,
	}

	misc.WriteJSONResponse(w, http.StatusCreated, resp)
}

func (c *ChatRoomHandler) leaveChatroom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roomId := ctx.Value(constants.RoomIDKey).(string)
	user := r.Header.Get("X-User-ID")
	if user == "" {
		http.Error(w, "Missing User ID", http.StatusUnauthorized)
		return
	}

	err := c.userManagementChatroomUseCase.Leave(roomId, user, ctx)

	if err != nil {
		misc.WriteJSONResponse(w, err.StatusCode, err.Messages)
	}

	misc.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Left chatroom successfully"})
}

func (c *ChatRoomHandler) joinChatroom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roomId := ctx.Value(constants.RoomIDKey).(string)
	user := r.Header.Get("X-User-ID")
	if user == "" {
		http.Error(w, "Missing User ID", http.StatusUnauthorized)
		return
	}
	err := c.userManagementChatroomUseCase.Join(roomId, user, ctx)

	if err != nil {
		misc.WriteJSONResponse(w, err.StatusCode, err.Messages)
	}

	misc.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Joined chatroom successfully"})
}

func (c *ChatRoomHandler) listChatroom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	listChatRoom, errResp := c.retrieveChatroomUseCase.ExecuteByFilter(ctx)

	if errResp != nil {
		misc.WriteJSONResponse(w, errResp.StatusCode, errResp.Messages)
	}

	var res []dto.ChatRoomDTO

	for _, v := range listChatRoom {
		res = append(res, dto.GetChatroomResponse(&v))
	}

	misc.WriteJSONResponse(w, http.StatusOK, res)
}

func (c *ChatRoomHandler) chatRoomDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	roomId := ctx.Value(constants.RoomIDKey).(string)

	res, errResp := c.retrieveChatroomUseCase.ExecuteById(roomId, ctx)

	if errResp != nil {
		misc.WriteJSONResponse(w, errResp.StatusCode, errResp.Messages)
	}

	misc.WriteJSONResponse(w, http.StatusOK, dto.GetChatroomResponse(res))
}

var ModuleChatRoomHandler = fx.Invoke(registerChatRoomHandlers)
