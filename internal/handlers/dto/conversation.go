package dto

type ConversationRequestDto struct {
	Destination string `json:"destination"`
	Type        string `json:"type"`
	Message     string `json:"message"`
}

type ConversationResponseDto struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}
