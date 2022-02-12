package models

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
