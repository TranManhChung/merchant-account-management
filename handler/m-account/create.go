package m_account

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Create, body: %v", request)

	json.NewEncoder(w).Encode(CreateResponse{
		Status: Success,
	})
}
