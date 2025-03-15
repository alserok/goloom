package models

type Config map[string]any

type UpdateConfigReq struct {
	Content string `json:"config"`
}

type Dir struct {
	Path    string          `json:"path"`
	Name    string          `json:"name"`
	Content map[string]bool `json:"content"`
}
