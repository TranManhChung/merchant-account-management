package m_account

import (
	"main.go/common/err"
	"main.go/db/m-account"
)

type CreateRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}

type ReadResponse struct {
	Status string                           `json:"status"`
	Error  *err.Error                       `json:"error,omitempty"`
	Data   *m_account.MerchantAccountEntity `json:"data"`
}

type UpdateRequest struct {
	MerchantID string `json:"merchant_id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
}

type UpdateResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}

type DeleteResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}
