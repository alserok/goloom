package models

type ServiceState struct {
	Service string `json:"service"`
	Status  int    `json:"status"`
}
