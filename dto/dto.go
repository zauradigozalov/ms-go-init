package dto

type UserRequest struct {
	UserName string `json:"userName"`
	Status   bool   `json:"status"`
}
