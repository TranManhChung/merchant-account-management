package m_account

import (
	"main.go/common/err"
	"main.go/model"
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
	Status string                       `json:"status"`
	Error  *err.Error                   `json:"error,omitempty"`
	Data   *model.MerchantAccountEntity `json:"data"`
}

type UpdateRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}

type DeleteRequest struct {
	Code string `json:"code"`
}

type DeleteResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}
