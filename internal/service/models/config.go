package models

type Config map[string]any

type UpdateConfigReq struct {
	Content string `json:"config"`
}

type Dir struct {
	Path    string    `json:"path"`
	Name    string    `json:"name"`
	Content []Content `json:"content"`
}

type Content struct {
	IsDir bool   `json:"is_dir"`
	Name  string `json:"name"`
	Size  string `json:"size"`
}
