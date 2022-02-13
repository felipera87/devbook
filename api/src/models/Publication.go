package models

import (
	"errors"
	"strings"
)

// Publication represents a publication made by an user
type Publication struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"authorId,omitempty"`
	AuthorNick string `json:"authorNick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreatedAt  uint64 `json:"createdAt,omitempty"`
}

// Prepare validates and formats all fields necessary for a new register
func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("title is required")
	}

	if publication.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)

}
