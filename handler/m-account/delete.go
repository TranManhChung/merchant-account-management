package m_account

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteRequest struct {
	Code     string `json:"code"`
}

type DeleteResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var request DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Delete, body: %v", request)

	json.NewEncoder(w).Encode(DeleteResponse{
		Status: Success,
	})
}
