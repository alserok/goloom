package models

type Config map[string]any

type UpdateConfigReq struct {
	Path   string `json:"path"`
	Config Config `json:"config"`
}
