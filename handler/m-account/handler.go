package m_account

import "net/http"

const (
	Success = "success"
	Failed = "failed"
)

type IHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Handler struct {}

type Error struct {
	Domain string `json:"domain"`
	Code int `json:"code"`
	Message string `json:"message"`
}