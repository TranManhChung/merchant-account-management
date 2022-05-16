package m_account

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

type UpdateResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	var request UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Update, body: %v", request)

	json.NewEncoder(w).Encode(UpdateResponse{
		Status: Success,
	})
}
