package m_account

import (
	"encoding/json"
	"fmt"
	"main.go/common/err"
	"net/http"
)

type UpdateRequest struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	DoB          string `json:"do_b"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender"`
}

type UpdateResponse struct {
	Status string `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
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
