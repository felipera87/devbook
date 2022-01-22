package models

// Password represents the necessary data for a password update
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
