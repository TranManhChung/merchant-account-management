package m_account

import (
	"main.go/common/err"
	"main.go/model"
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

type ReadResponse struct {
	Status string                      `json:"status"`
	Error  *err.Error                      `json:"error,omitempty"`
	Data   *model.MerchantMemberEntity `json:"data"`
}

type DeleteRequest struct {
	Email string `json:"email"`
}

type DeleteResponse struct {
	Status string `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}
