package model

type TypeInfo struct {
	TypeID string `json:"typeID"`
	Name   string `json:"name"`
}

type NameInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
