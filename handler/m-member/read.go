package m_account

import (
	"encoding/json"
	"fmt"
	"main.go/common/err"
	"main.go/model"
	"net/http"
)

type ReadResponse struct {
	Status string                      `json:"status"`
	Error  *err.Error                      `json:"error,omitempty"`
	Data   *model.MerchantMemberEntity `json:"data"`
}

func (h Handler) Read(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("email")

	fmt.Printf("Read, body: %v", id)

	json.NewEncoder(w).Encode(ReadResponse{
		Status: Success,
	})
}
