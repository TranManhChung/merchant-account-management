package m_account

import (
	"encoding/json"
	"fmt"
	"main.go/common/err"
	"net/http"
)

type CreateRequest struct {
	Email        string `json:"email"`
	MerchantCode string `json:"merchant_code"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	DoB          string `json:"do_b"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender"`
}

type CreateResponse struct {
	Status string `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
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
