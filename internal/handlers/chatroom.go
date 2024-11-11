package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go-live-chat/internal/handlers/dto"
	"go-live-chat/internal/misc"
	"go.uber.org/fx"
	"net/http"
)

type ChatRoomHandler struct {
	createChatroomUseCase CreateChatroomUseCase
}

func NewChatRoomHandler(createChatroomUseCase CreateChatroomUseCase) *ChatRoomHandler {
	return &ChatRoomHandler{
		createChatroomUseCase: createChatroomUseCase,
	}
}

func registerChatRoomHandlers(c *ChatRoomHandler, h *Handler) {
	h.Runner.Group(func(r chi.Router) {
		//r.Use(tokenValidationMiddleware)
		r.Route("/api/chatrooms", func(r chi.Router) {
			r.Post("/", c.createChatroom)

			r.Route("/{roomId}", func(r chi.Router) {

				r.Use(chatroomCtx)

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
	_, err := w.Write([]byte("Leave Chatroom"))
	if err != nil {
		return
	}
}

func (c *ChatRoomHandler) joinChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Join Chatroom"))
	if err != nil {
		return
	}
}

var ModuleChatRoomHandler = fx.Invoke(registerChatRoomHandlers)
