package m_account

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ReadResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
}

func (h Handler) Read(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Printf("Read, body: %v", id)

	json.NewEncoder(w).Encode(ReadResponse{
		Status: Success,
	})
}