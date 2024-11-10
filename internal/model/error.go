package model

type (
	Error struct {
		StatusCode int
		Messages   []string
	}
)
