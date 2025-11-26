package models

type FileData struct {
	Version string      `json:"version"`
	Data    []Workspace `json:"data"`
}
