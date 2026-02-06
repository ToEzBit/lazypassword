package models

type FileData struct {
	Version      string      `json:"version"`
	Salt         string      `json:"salt"`
	PasswordHash string      `json:"password_hash,omitempty"`
	Data         []Workspace `json:"data"`
}
