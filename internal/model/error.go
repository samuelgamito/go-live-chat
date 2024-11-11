package model

type (
	Error struct {
		StatusCode int
		Messages   []string
	}
)

func (e Error) Error() string {
	return e.Messages[0]
}
